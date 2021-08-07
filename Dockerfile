FROM golang:alpine3.14

WORKDIR /

COPY . .


RUN go mod download && go build -o bin/words main.go

CMD ["./bin/words"]