# Build stage
FROM golang:1.22.2-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go 

# RUN stage
FROM alpine:3.19.1 

WORKDIR /app

COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080

CMD [ "/app/main"]



