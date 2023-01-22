FROM python:3.8

WORKDIR /code

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY chromedriver.exe *.py ./src/

ENTRYPOINT [ "python3", "-u", "./src/main.py" ]

