FROM golang:1.18.3-alpine as builder

WORKDIR /app

COPY go.mod go.sum main.go /app/

RUN GOOS=linux go build

FROM alpine:latest

COPY --from=builder /app/go-service-1 /go-service-1

CMD ["/go-service-1"]
