FROM python:3.8

WORKDIR /code

RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt update -y && apt upgrade -y
RUN apt install ./google-chrome-stable_current_amd64.deb -y

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY chromedriver.exe *.py ./src/
EXPOSE 8083

ENTRYPOINT [ "python3", "-u", "./src/main.py" ]

