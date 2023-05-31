package user

import (
	"context"
	"net/http"
	"time"

	"github.com/caiomp87/lost-and-found/src/modules/user/dto"
	"github.com/caiomp87/lost-and-found/src/modules/user/services"
	"github.com/gofiber/fiber/v2"
)

type IUserHandler interface {
	Create(c *fiber.Ctx) error
}

type UserHandler struct {
	service services.IUserService
}

func NewUserHandler(service services.IUserService) IUserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	userDto := new(dto.CreateUserRequest)

	if err := c.JSON(&userDto); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "could not process entity: " + err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := h.service.Create(ctx, userDto); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create user: " + err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
	})
}
