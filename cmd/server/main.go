package main

import (
	"fmt"

	sqlx "github.com/jmoiron/sqlx"

	http "github.com/ImKairat-Golang-Lab/users-service/internal/adapters/http"
	logger "github.com/ImKairat-Golang-Lab/users-service/internal/adapters/logger"
	psql_repo "github.com/ImKairat-Golang-Lab/users-service/internal/adapters/repository/psql"
	services "github.com/ImKairat-Golang-Lab/users-service/internal/domain/services"
	utils "github.com/ImKairat-Golang-Lab/users-service/internal/domain/utils"
)

func main() {
	logger := logger.NewZapLogger()

	dsn := "postgres://user:password@localhost:5432/mydb?sslmode=disable"
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		logger.Warn(fmt.Sprintf("failed to connect to DB: %v", err), map[string]any{})
	}

	rc := utils.RealClock{}

	repo := psql_repo.NewPostgresUserRepository(db, logger)
	service := services.NewUserService(repo, rc)
	handler := http.NewUserHandler(service, logger)

	server := http.NewHTTPServer(*handler)
	err = server.Start("127.0.0.1", 8080)
	if err != nil {
		logger.Warn(
			fmt.Sprintf("failed to connect %v", err),
			map[string]any{},
		)
	}
}
