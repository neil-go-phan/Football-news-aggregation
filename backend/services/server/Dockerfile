# Build stage
FROM golang:1.20-alpine3.16 AS builder
LABEL author="Neil GoldenOwl golang intern"
WORKDIR /app
COPY . . 
RUN go build -o main main.go

EXPOSE 8080


# Run stage
FROM alpine:3.16

ENV TZ=Asia/Ho_Chi_Minh

WORKDIR /app
COPY .env .
COPY --from=builder /app/main .

RUN apk add --no-cache tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime

COPY wait-for.sh .

CMD [ "/app/main" ]