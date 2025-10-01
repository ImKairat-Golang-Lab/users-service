package ports

import domain "github.com/ImKairat-Golang-Lab/user-service/internal/models"

type User = domain.User

type UserRepository interface {
	Save(user User) error
	FindById(id string) (User, error)
	FindByEmail(email string) (User, error)
	Update(user User) error
	Delete(id string) error
}
