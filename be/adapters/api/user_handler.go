package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/peesaphanthavong/core/usecases"
	"github.com/peesaphanthavong/models"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetAllUser(ctx *fiber.Ctx) error {

	result, err := h.userUseCase.InquryALLUser(ctx)
	if err != nil {
		return models.NewResponse("500", "internal error", err.Error()).SendResponse(ctx, http.StatusInternalServerError)
	}
	return models.NewResponse("0000", "success", result).SendResponse(ctx, http.StatusOK)

}


func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var request usecases.CreateUserRequest
	if err := ctx.BodyParser(&request); err != nil {
		return models.NewResponse("400", "bad request", err.Error()).SendResponse(ctx, http.StatusBadRequest)
	}

	if err := h.userUseCase.CreateUser(ctx, request); err != nil {
		return models.NewResponse("500", "internal error", err.Error()).SendResponse(ctx, http.StatusInternalServerError)
	}

	return models.NewResponse("0000", "success", nil).SendResponse(ctx, http.StatusCreated)
}