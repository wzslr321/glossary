FROM python:3.9.6-alpine3.14

WORKDIR /usr/src/scrapper

COPY requirements.txt ./

RUN apk add build-base

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD ["python","main.py"]
