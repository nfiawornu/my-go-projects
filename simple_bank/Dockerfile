# Build stage
FROM golang:1.24.6-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run stage
FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8082
CMD [ "/app/main" ]