FROM golang:1.26.2-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
