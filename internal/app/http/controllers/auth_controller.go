package controllers

import (
	"github.com/asynched/gist-backend/internal/app/schemas"
	"github.com/asynched/gist-backend/internal/database/repositories"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	userRepository *repositories.UserRepository
}

func (controller *AuthController) Register(ctx *fiber.Ctx) error {
	data := schemas.CreateUserSchema{}

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	if valid, errors := data.IsValid(); !valid {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
			"errors":  errors,
		})
	}

	err := controller.userRepository.Create(repositories.CreateUserInput{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func NewAuthController(userRepository *repositories.UserRepository) *AuthController {
	return &AuthController{
		userRepository: userRepository,
	}
}
