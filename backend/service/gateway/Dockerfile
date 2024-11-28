FROM golang:1.21-alpine as builder

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o gateway main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/gateway .
EXPOSE 8080
CMD ["./gateway"]