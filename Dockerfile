FROM golang:buster

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o cmd/main.go

RUN ["/app/main"]
