package http

import (
	"github.com/asynched/gist-backend/internal/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	statusController := controllers.NewStatusController()

	app.Get("/status", statusController.GetStatus)
}
