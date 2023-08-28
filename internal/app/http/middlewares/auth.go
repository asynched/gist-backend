package middlewares

import (
	"github.com/asynched/gist-backend/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	jwtService *services.JwtService
}

func (middleware *AuthMiddleware) Handle(ctx *fiber.Ctx) error {
	auth := ctx.Get("Authorization")

	if auth == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Authorization header is required",
		})
	}

	token := auth[7:]

	if token == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Token is required",
		})
	}

	claims := jwt.MapClaims{}

	_, err := middleware.jwtService.ValidateToken(token, claims)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	ctx.Locals("userId", int64(claims["userId"].(float64)))

	return ctx.Next()
}

func NewAuthMiddleware(jwtService *services.JwtService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}
