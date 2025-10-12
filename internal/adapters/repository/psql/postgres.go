package psql

import (
	"context"

	models "github.com/ImKairat-Golang-Lab/users-service/internal/adapters/repository/psql/models"
	entities "github.com/ImKairat-Golang-Lab/users-service/internal/domain/entities"
	ports "github.com/ImKairat-Golang-Lab/users-service/internal/ports"
	sqlx "github.com/jmoiron/sqlx"
)

type User = entities.User
type UserModel = models.UserModel

// type repo = ports.UserRepository

type PostgresUserRepository struct {
	db     *sqlx.DB
	logger ports.Logger
}

func NewPostgresUserRepository(db *sqlx.DB, logger ports.Logger) *PostgresUserRepository {
	return &PostgresUserRepository{
		db:     db,
		logger: logger,
	}
}

func (ur *PostgresUserRepository) Save(ctx context.Context, user User) error {
	component := "PostgresUserRepository/Save"
	// TODO: добавить логику для получения названия таблицы (users)
	var user_model = UserModel{
		Id:           user.Id,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Login:        user.Login,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	query := `INSERT INTO users (id, email, password_hash, login, created_at, updated_at)
			  VALUES (:id, :email, :password_hash, :login, :created_at, :updated_at)`

	_, err := ur.db.NamedExec(query, user_model)
	if err != nil {
		ur.logger.Warn(err.Error(), map[string]any{
			"component": component,
			"user_id":   user_model.Id,
		})
		return err
	}

	return nil
}
