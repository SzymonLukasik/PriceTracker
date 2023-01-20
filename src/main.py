from datetime import datetime

import selenium.common.exceptions
from google import protobuf
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from google.protobuf.timestamp_pb2 import Timestamp
import Products_pb2_grpc as pygrpc
import Products_pb2 as pb
import schedule
import grpc

DRIVER_PATH = "./chromedriver.exe"
options = Options()
# options.headless = True
options.add_argument("--window-size=1920,1200")


def scrape_media_markt(addr):
    price = 0
    tries = 0
    while price == 0 and tries < 3:
        tries += 1
        try:
            """element = WebDriverWait(driver, 10).until(
                EC.presence_of_element_located((By.CLASS_NAME, "whole"))
            )"""
            driver = webdriver.Chrome(options=options, executable_path=DRIVER_PATH)
            driver.get(addr)
            price_cands = driver.find_elements(By.CLASS_NAME, "is-big")
            for cand in price_cands:
                if cand.text != '':
                    print(cand.text)
                    price_raw = cand.text.replace(',', '').replace('-', '').split('\n', 1)
                    price += int(price_raw[0]) * 100
                    if len(price_raw) > 1 and len(price_raw[1]) > 0:
                        price += int(price_raw[1])
                    break
        except selenium.common.exceptions.TimeoutException as e:
            print((e.msg, tries, "retrying 3 times"))
        finally:
            driver.quit()
    return price


def scrape_euro(addr):
    price = 0
    tries = 0
    class_name = "selenium-price-normal"
    while price == 0 and tries < 3:
        print(('already tried times:', tries))
        tries += 1
        try:
            driver = webdriver.Chrome(options=options, executable_path=DRIVER_PATH)
            driver.get(addr)
            """print('waiting')
            element = WebDriverWait(driver, 10).until(
                EC.presence_of_element_located((By.CLASS_NAME, class_name))
            )
            print('waited')"""
            price_cands = driver.find_elements(By.CLASS_NAME, class_name)
            for cand in price_cands:
                if cand.text != '':
                    price_raw = cand.text[:len(cand.text) - 3].replace(' ', '')
                    mult = 100
                    if price_raw.find(','):
                        mult = 1
                        price_raw = price_raw.replace(',', '')
                    price += int(price_raw) * mult
                    break
        except selenium.common.exceptions.TimeoutException as e:
            print((e.msg, tries, "retrying 3 times"))
        finally:
            driver.quit()
    return price


def scrape():
    channel = grpc.insecure_channel('localhost:8083')
    stub = pygrpc.ProductsStub(channel)
    prods = stub.GetAllProducts(protobuf.empty_pb2.Empty())
    if prods is None:
        print("unable to get products")
        return
    prods = prods.productsList
    for p in prods:
        dp = pb.DatePrice()
        dp.ts.FromDatetime(datetime.now())
        price = -1
        if p.shop == "Euro":
            price = scrape_euro(p.url)
            print(price)
        elif p.shop == "MediaMarkt":
            price = scrape_media_markt(p.url)
            print(price)
        else:
            print("unknown shop")
            return
        dp.price = price
        request = pb.ProductNewPrice(price=dp, product=p)
        stub.AddNewPrice(request)


if __name__ == '__main__':
    schedule.every(1).minutes.do(scrape)
    schedule.run_all()
