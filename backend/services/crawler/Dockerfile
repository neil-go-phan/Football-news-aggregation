# Build stage
FROM golang:1.20-alpine3.16 AS builder
LABEL author="Neil GoldenOwl golang intern"
WORKDIR /app
COPY . . 
RUN go build -o main main.go

EXPOSE 8000

# Run stage
FROM alpine:3.16

ENV TZ=Asia/Ho_Chi_Minh
RUN apk update && apk upgrade && \
    apk add --no-cache \
        chromium \
        nss \
        freetype \
        freetype-dev \
        harfbuzz \
        ca-certificates \
        ttf-freefont
       
WORKDIR /app
COPY configs/*.json /app/configs/
COPY .env .
COPY db/migrations ./db/migrations
COPY --from=builder /app/main .

RUN apk add --no-cache tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime

CMD [ "/app/main" ]