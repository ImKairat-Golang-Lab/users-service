package ports

import (
	"context"
)

type UserService interface {
	Register(ctx context.Context, email, password, login string) error
}
