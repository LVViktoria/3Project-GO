FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o weather \
    ./cmd/weather

FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/weather ./weather

COPY --from=builder /app/data ./data

ENTRYPOINT ["./weather"]
