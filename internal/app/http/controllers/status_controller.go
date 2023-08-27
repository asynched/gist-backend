package controllers

import "github.com/gofiber/fiber/v2"

type StatusController struct {
}

func (controller *StatusController) GetStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func NewStatusController() *StatusController {
	return &StatusController{}
}
