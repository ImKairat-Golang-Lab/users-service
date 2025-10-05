package entities

import (
	"time"
)

// Логическая единица, с которой связано работа бизнес-процесса:
type User struct {
	Id           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Login         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
