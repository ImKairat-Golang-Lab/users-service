package psql

import (
	"context"

	"github.com/ImKairat-Golang-Lab/users-service/internal/domain/entities"
	// "github.com/ImKairat-Golang-Lab/users-service/internal/ports"
	"github.com/ImKairat-Golang-Lab/users-service/internal/adapters/repository/psql/models"
	"github.com/jmoiron/sqlx"
)

type User = entities.User
type UserModel = models.UserModel

// type repo = ports.UserRepository

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(ctx context.Context, user User) error {
	// TODO: добавить логику для получения названия таблицы (users)
	var user_model UserModel = UserModel{
		Id: user.Id,
		Email: user.Email,
		PasswordHash: user.PasswordHash,
		Login: user.Login,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	query := `INSERT INTO users (id, email, password_hash, login, created_at, updated_at)
			  VALUES (:id, :email, :password_hash, :login, :created_at, :updated_at)`

	_, err := r.db.NamedExec(query, user_model)
	if err != nil {
		return err
	}
	
	return nil
}
