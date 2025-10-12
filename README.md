# user-service

## Sync all packages

```bash
go mod tidy
```

## Install linter

```bash
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
```

## Install testing libs

```bash
go get github.com/golang/mock/gomock@v1.6.0
```

```bash
go install github.com/golang/mock/mockgen@latest
```

## Generate mocks

```bash
mockgen -source=internal/ports/user_repository.go -destination=internal/mocks/mock_user_repository.go -package=mocks
```

## Run linter

```bash
golangci-lint run ./...
```

## Run tests

```bash
go test ./... -v -cover
```
