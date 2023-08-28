package controllers

import (
	"github.com/asynched/gist-backend/internal/database/repositories"
	"github.com/gofiber/fiber/v2"
)

type GistsController struct {
	gistRepository *repositories.GistRepository
}

func (controller *GistsController) GetGists(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(int64)

	gists, err := controller.gistRepository.FindGistsByUserId(repositories.FindGistsByUserIdInput{
		UserId: userId,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get gists",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(gists)
}

func NewGistsController(gistRepository *repositories.GistRepository) *GistsController {
	return &GistsController{
		gistRepository: gistRepository,
	}
}
