FROM golang:1.26.3 AS builder

WORKDIR /cmd

COPY . .

RUN go build -o pitsinkr .