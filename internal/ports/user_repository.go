package ports

import (
	"context"
	"github.com/ImKairat-Golang-Lab/users-service/internal/domain/entities"
)

// Для читабельности добавил как алиас:
type User = entities.User

// Мост между доменом и адаптерами.
type UserRepository interface {
	Save(ctx context.Context, user User) error
	FindById(ctx context.Context, id string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
