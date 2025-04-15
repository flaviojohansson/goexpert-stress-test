FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o goexpert-stress-test .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/goexpert-stress-test .

ENTRYPOINT ["./goexpert-stress-test"]