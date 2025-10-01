package inmemory

import (
	"github.com/ImKairat-Golang-Lab/users-service/internal/models"
	"github.com/ImKairat-Golang-Lab/users-service/internal/ports"
)

type User = models.User

type MemoryUserRepository struct {
	repo ports.UserRepository
}


func NewMemoryUserRepository(repo *ports.UserRepository) *MemoryUserRepository {
	return &MemoryUserRepository{repo: *repo}
}
