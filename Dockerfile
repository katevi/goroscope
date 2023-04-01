# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum main.go ./
COPY ./internal ./internal

RUN go mod tidy
RUN go mod download

RUN go build -v -o /goroscope .

CMD [ "/goroscope" ]