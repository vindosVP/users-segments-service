FROM golang:1.20.0

WORKDIR /usr/src/app

COPY ./src ./src
WORKDIR /usr/src/app/src
RUN go mod tidy
