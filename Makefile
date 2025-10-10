
run:
	@go run ./cmd/server/main.go


fmt:
	@gofmt -w .


up:
	@docker compose up -d

down:
	@docker compose down

lint:
	@golangci-lint run ./...
