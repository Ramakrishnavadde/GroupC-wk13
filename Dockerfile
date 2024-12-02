FROM golang:1.20 AS build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o toronto-time-api .

FROM alpine:latest

RUN apk add --no-cache mysql-client

WORKDIR /root
COPY --from=build /app/toronto-time-api /usr/local/bin/toronto-time-api

CMD ["toronto-time-api"]
