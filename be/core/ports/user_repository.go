package ports

import (
	"context"

	"github.com/peesaphanthavong/models"
	"github.com/peesaphanthavong/models/response/user"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]user.UserResponse, error)
	CreateUser(ctx context.Context, user *models.User) error
	// GetUserByUserId(ctx context.Context, id string) (*models.User, error)
}
