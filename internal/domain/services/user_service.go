package services

import (
	"context"
	"time"

	"github.com/ImKairat-Golang-Lab/users-service/internal/domain/entities"
	"github.com/ImKairat-Golang-Lab/users-service/internal/ports"
)

type User = entities.User

// Для DI (Dependencies injection)
type UserService struct {
	repo ports.UserRepository
}

//
func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

//
func (us *UserService) Register(ctx context.Context, email, password, login string) error {
	id := "" // Необходимо реализовать функцию для генерации уникального ID

	user := User{
		Id:           id,
		Email:        email,
		PasswordHash: password,
		Login:         login,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return us.repo.Save(ctx, user)
}

func Login(ctx context.Context, email, password string) {}

func GetProfile(ctx context.Context, id string) {}

func UpdateProfile(ctx context.Context, id, name string) {}

func DeleteProfile(ctx context.Context, id string) {}
