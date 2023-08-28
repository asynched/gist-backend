package controllers

import (
	"github.com/asynched/gist-backend/internal/app/schemas"
	"github.com/asynched/gist-backend/internal/database/repositories"
	"github.com/asynched/gist-backend/internal/utils"
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

func (controller *GistsController) CreateGist(ctx *fiber.Ctx) error {
	data := schemas.CreateGistSchema{}

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create gist",
			"error":   err.Error(),
		})
	}

	if valid, errors := data.IsValid(); !valid {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation error",
			"errors":  errors,
		})
	}

	userId := ctx.Locals("userId").(int64)

	gist, err := controller.gistRepository.Create(repositories.CreateGistInput{
		UserId:      userId,
		Title:       data.Title,
		Description: data.Description,
		Files: utils.Map(data.Files, func(file schemas.CreateFileSchema) repositories.CreateFileInput {
			return repositories.CreateFileInput{
				Filename: file.Filename,
				Content:  file.Content,
			}
		}),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create gist",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(gist)
}

func NewGistsController(gistRepository *repositories.GistRepository) *GistsController {
	return &GistsController{
		gistRepository: gistRepository,
	}
}
