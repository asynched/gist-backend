package controllers

import (
	"strconv"

	"github.com/asynched/gist-backend/internal/app/schemas"
	"github.com/asynched/gist-backend/internal/database/repositories"
	"github.com/gofiber/fiber/v2"
)

type FilesController struct {
	filesRepository *repositories.FileRepository
	gistRepository  *repositories.GistRepository
}

func (controller *FilesController) GetFiles(ctx *fiber.Ctx) error {
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
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Could not get gist",
			"error":   err.Error(),
		})
	}

	if gist.UserId != userId {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not authorized to view this gist",
		})
	}

	files, err := controller.filesRepository.FindFilesByGistId(repositories.FindFilesByGistIdInput{
		GistId: int64(gistId),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get files",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(files)
}

func (controller *FilesController) GetFile(ctx *fiber.Ctx) error {
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
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Could not get gist",
			"error":   err.Error(),
		})
	}

	if gist.UserId != userId {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not authorized to view this gist",
		})
	}

	fileId, err := strconv.Atoi(ctx.Params("fileId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid file id",
			"error":   err.Error(),
		})
	}

	file, err := controller.filesRepository.FindFileById(repositories.FindFileByIdInput{
		FileId: int64(fileId),
	})

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Could not get file",
			"error":   err.Error(),
		})
	}

	if file.GistId != int64(gistId) {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not authorized to view this file",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(file)
}

func (controller *FilesController) UpdateFile(ctx *fiber.Ctx) error {
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
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Could not get gist",
			"error":   err.Error(),
		})
	}

	if gist.UserId != userId {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not authorized to update this gist",
		})
	}

	fileId, err := strconv.Atoi(ctx.Params("fileId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid file id",
			"error":   err.Error(),
		})
	}

	data := schemas.UpdateFileSchema{}

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update file",
			"error":   err.Error(),
		})
	}

	file, err := controller.filesRepository.FindFileById(repositories.FindFileByIdInput{
		FileId: int64(fileId),
	})

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Could not get file",
			"error":   err.Error(),
		})
	}

	if file.GistId != int64(gistId) {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not authorized to update this file",
		})
	}

	file, err = controller.filesRepository.UpdateFile(repositories.UpdateFileInput{
		FileId:   int64(fileId),
		Filename: data.Filename,
		Content:  data.Content,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update file",
			"error":   err.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(file)
}

func (controller *FilesController) CreateFile(ctx *fiber.Ctx) error {
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
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Could not get gist",
			"error":   err.Error(),
		})
	}

	if gist.UserId != userId {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not authorized to create a file for this gist",
		})
	}

	data := schemas.CreateFileSchema{}

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create file",
			"error":   err.Error(),
		})
	}

	file, err := controller.filesRepository.CreateFile(repositories.CreateFileInput{
		GistId:   int64(gistId),
		Filename: data.Filename,
		Content:  data.Content,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create file",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(file)
}

func (controller *FilesController) DeleteFile(ctx *fiber.Ctx) error {
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
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Could not get gist",
			"error":   err.Error(),
		})
	}

	if gist.UserId != userId {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not authorized to delete this gist",
		})
	}

	fileId, err := strconv.Atoi(ctx.Params("fileId"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid file id",
			"error":   err.Error(),
		})
	}

	err = controller.filesRepository.DeleteFile(repositories.DeleteFileInput{
		FileId: int64(fileId),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete file",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}

func NewFilesController(filesRepository *repositories.FileRepository, gistRepository *repositories.GistRepository) *FilesController {
	return &FilesController{
		filesRepository: filesRepository,
		gistRepository:  gistRepository,
	}
}
