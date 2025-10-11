
.PHONY: run fmt up down lint generate-mocks test prep2commit
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

generate-mocks:
	@mockgen -source=internal/ports/user_repository.go -destination=internal/mocks/mock_user_repository.go -package=mocks
	@mockgen -source=internal/ports/user_service.go -destination=internal/mocks/mock_user_service.go -package=mocks
	@mockgen -source=internal/domain/utils/clock.go -destination=internal/mocks/mock_clock.go -package=mocks

test:
	@go test ./...

prep2commit: lint test
