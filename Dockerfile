FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o go-order-api

EXPOSE 8080

CMD ./go-order-api