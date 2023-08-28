package controllers

import (
	"github.com/asynched/gist-backend/internal/app/schemas"
	"github.com/asynched/gist-backend/internal/database/repositories"
	"github.com/asynched/gist-backend/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	userRepository *repositories.UserRepository
	jwtService     *services.JwtService
}

func (controller *AuthController) SignUp(ctx *fiber.Ctx) error {
	data := schemas.SignUpSchema{}

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
		Username: data.Username,
		Password: data.Password,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (controller *AuthController) SignIn(ctx *fiber.Ctx) error {
	data := schemas.SignInSchema{}

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

	user, err := controller.userRepository.FindUserByUsername(repositories.FindUserByUsernameInput{
		Username: data.Username,
	})

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	if user.Password != data.Password {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	claims := jwt.MapClaims{
		"userId": user.UserId,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": controller.jwtService.GenerateToken(claims),
	})
}

func NewAuthController(userRepository *repositories.UserRepository, jwtService *services.JwtService) *AuthController {
	return &AuthController{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}
