package controllers

import (
	"strconv"

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

	gists, err := controller.gistRepository.FindGists(repositories.FindGistsInput{
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

func (controller *GistsController) GetGist(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(int64)
	gistId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid gist id",
			"error":   err.Error(),
		})
	}

	gist, err := controller.gistRepository.FindGistById(repositories.FindGistByIdInput{
		GistId: int64(gistId),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get gist",
			"error":   err.Error(),
		})
	}

	if gist.UserId != userId {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not authorized to view this gist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(gist)
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

	gist, err := controller.gistRepository.CreateGist(repositories.CreateGistInput{
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

func (controller *GistsController) DeleteGist(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(int64)
	gistId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid gist id",
			"error":   err.Error(),
		})
	}

	gist, err := controller.gistRepository.FindGistById(repositories.FindGistByIdInput{
		GistId: int64(gistId),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get gist",
			"error":   err.Error(),
		})
	}

	if gist.UserId != userId {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not authorized to delete this gist",
		})
	}

	err = controller.gistRepository.DeleteGist(repositories.DeleteGistInput{
		GistId: int64(gistId),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete gist",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

func (controller *GistsController) UpdateGist(ctx *fiber.Ctx) error {
	data := schemas.UpdateGistSchema{}

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update gist",
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
	gistId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid gist id",
			"error":   err.Error(),
		})
	}

	gist, err := controller.gistRepository.FindGistById(repositories.FindGistByIdInput{
		GistId: int64(gistId),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get gist",
			"error":   err.Error(),
		})
	}

	if gist.UserId != userId {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not authorized to update this gist",
		})
	}

	gist, err = controller.gistRepository.UpdateGist(repositories.UpdateGistInput{
		GistId:      int64(gistId),
		Title:       data.Title,
		Description: data.Description,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update gist",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(gist)
}

func NewGistsController(gistRepository *repositories.GistRepository) *GistsController {
	return &GistsController{
		gistRepository: gistRepository,
	}
}
