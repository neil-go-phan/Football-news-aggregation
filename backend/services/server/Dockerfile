# Build stage
FROM golang:1.20-alpine3.16 AS builder
LABEL author="Neil GoldenOwl golang intern"
WORKDIR /app
COPY . . 
RUN go build -o main main.go

EXPOSE 8080

# Run stage
FROM alpine:3.16

WORKDIR /app
COPY configs/*.json /app/configs/
COPY app.env .
COPY --from=builder /app/main .
COPY wait-for.sh .

CMD [ "/app/main" ]