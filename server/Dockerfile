FROM golang:alpine3.14

WORKDIR /server

RUN apk add git gcc libc-dev \
    &&  go get -u -v github.com/cosmtrek/air \
    && apk add build-base

COPY .. .

CMD ["air"]