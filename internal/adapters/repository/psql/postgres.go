package psql

import (
	"github.com/ImKairat-Golang-Lab/user-service/internal/models"
	// "github.com/ImKairat-Golang-Lab/user-service/internal/ports"
	"github.com/jmoiron/sqlx"
)

type User = models.User

// type repo = ports.UserRepository

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(user User) error {
	return nil
}
