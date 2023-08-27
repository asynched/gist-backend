package main

import (
	"log"

	"github.com/asynched/gist-backend/internal/app/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(requestid.New())
	http.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
