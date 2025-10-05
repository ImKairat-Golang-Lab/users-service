package psql

import (
	"context"

	"github.com/ImKairat-Golang-Lab/users-service/internal/domain/entities"
	// "github.com/ImKairat-Golang-Lab/users-service/internal/ports"
	"github.com/jmoiron/sqlx"
)

type User = entities.User

// type repo = ports.UserRepository

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(ctx context.Context, user User) error {
	return nil
}
