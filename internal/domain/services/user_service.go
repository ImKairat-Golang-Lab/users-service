package services

import (
	"context"
	"errors"
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
	if us.repo == nil {
		return errors.New("UserRepository doesn't initialized")
	}
	// TODO:Необходимо реализовать функционал для генерации уникального ID для пользователья 
	id := ""

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
