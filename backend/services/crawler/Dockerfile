# Build stage
FROM golang:1.20-alpine3.16 AS builder
LABEL author="Neil GoldenOwl golang intern"
WORKDIR /app
COPY . . 
RUN go build -o main main.go

EXPOSE 8000

# Run stage
FROM alpine:3.16

WORKDIR /app
COPY configs/*.json /app/configs/
COPY .env .
COPY --from=builder /app/main .

CMD [ "/app/main" ]