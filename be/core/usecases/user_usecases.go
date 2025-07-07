package usecases

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peesaphanthavong/core/ports"
	"github.com/peesaphanthavong/models"
	"github.com/peesaphanthavong/models/response/user"
)

type UserUseCaser interface {
	InquryALLUser(ctx *fiber.Ctx) ([]models.User, error)
	CreateUser(ctx *fiber.Ctx, request CreateUserRequest) error
}

type UserUseCase struct {
	userRepository ports.UserRepository
}

func NewUserUseCase(userRepository ports.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (uc *UserUseCase) InquryALLUser(ctx *fiber.Ctx) ([]user.UserResponse, error) {
	users, err := uc.userRepository.GetAllUser(ctx.Context())
	if err != nil {
		return nil, err
	}
	return users, nil
}


func (uc *UserUseCase) CreateUser(ctx *fiber.Ctx, request CreateUserRequest) error {
	user := models.User{
		UserName:  request.UserName,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}

	if err := uc.userRepository.CreateUser(ctx.Context(), &user); err != nil {
		return err
	}
	return nil
}