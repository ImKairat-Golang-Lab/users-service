package domain

import (
	"time"

	"github.com/ImKairat-Golang-Lab/users-service/internal/models"
	"github.com/ImKairat-Golang-Lab/users-service/internal/ports"
)

type User = models.User

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

//
func (us *UserService) Register(email, password, login string) {
	user := User{
		Id:           "",
		Email:        email,
		PasswordHash: password,
		Login:         login,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	us.repo.Save(user)
}

func Login(email, password string) {}

func GetProfile(id string) {}

func UpdateProfile(id, name string) {}

func DeleteProfile(id string) {}
