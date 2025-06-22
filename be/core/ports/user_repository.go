package ports

import (
	"context"

	"github.com/peesaphanthavong/models"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	// GetUserByUserId(ctx context.Context, id string) (*models.User, error)
}
