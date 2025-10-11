# user-service

Install linter:

```bash
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
```

```bash
go get github.com/golang/mock/gomock@v1.6.0
```

```bash
go install github.com/golang/mock/mockgen@latest
```

```bash
mockgen -source=internal/ports/user_repository.go -destination=internal/mocks/mock_user_repository.go -package=mocks
```
