FROM golang:1.23.4 AS builder
WORKDIR /app

COPY . .
ARG CGO_ENABLED=0 \
    GOOS=linux

RUN go mod tidy && go build -o ./bin/app ./cmd/main.go

FROM alpine:3.21.3
COPY --from=builder /app/bin/app /

EXPOSE 8080
ENTRYPOINT ["/app"]