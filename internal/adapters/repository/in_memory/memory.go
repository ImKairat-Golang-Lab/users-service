package inmemory

import (
	"github.com/ImKairat-Golang-Lab/users-service/internal/domain/entities"
	"github.com/ImKairat-Golang-Lab/users-service/internal/ports"
)

type User = entities.User

type MemoryUserRepository struct {
	repo ports.UserRepository
}

func NewMemoryUserRepository(repo *ports.UserRepository) *MemoryUserRepository {
	return &MemoryUserRepository{repo: *repo}
}
