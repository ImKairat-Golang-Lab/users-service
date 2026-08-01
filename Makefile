
.PHONY: run fmt up down lint generate-mocks test prep2commit init
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

gen-mocks:
	@mockgen -source=internal/ports/user_repository.go -destination=internal/mocks/mock_user_repository.go -package=mocks
	@mockgen -source=internal/ports/user_service.go -destination=internal/mocks/mock_user_service.go -package=mocks
	@mockgen -source=internal/domain/utils/clock.go -destination=internal/mocks/mock_clock.go -package=mocks

test:
	@go test -parallel 4 ./...  -cover

test-cover:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out -o coverage.html

prep2commit: fmt lint test

init:
	@go mod tidy
	@go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	@go get github.com/golang/mock/gomock@v1.6.0
	@go install github.com/golang/mock/mockgen@latest

gen-graph:
	@goda graph ./... | dot -Tsvg -o graph.svg
