FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN ls -a
RUN go build -o main ./cmd/main.go

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]