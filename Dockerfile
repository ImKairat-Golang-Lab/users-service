FROM golang:1.25.1 AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY vendor ./vendor
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o main ./cmd/server/

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
