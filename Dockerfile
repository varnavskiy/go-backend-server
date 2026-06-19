#Stage 1. Building
FROM golang:1.26-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o backendServer ./cmd/server/

#Stage 2. Running
FROM scratch

WORKDIR /usr/src/app

COPY .system.env .
COPY --from=builder /usr/src/app/backendServer .

CMD ["./backendServer"]
