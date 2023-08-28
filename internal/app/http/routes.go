package http

import (
	"github.com/asynched/gist-backend/internal/app/http/controllers"
	"github.com/asynched/gist-backend/internal/app/http/middlewares"
	"github.com/asynched/gist-backend/internal/database"
	"github.com/asynched/gist-backend/internal/database/repositories"
	"github.com/asynched/gist-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Database
	db := database.CreateClient()

	// Services and repositories
	userRepository := repositories.NewUserRepository(db)
	gistRepository := repositories.NewGistRepository(db)
	fileRepository := repositories.NewFileRepository(db)
	jwtService := services.NewJwtService()

	// Middlewares
	authMiddleware := middlewares.NewAuthMiddleware(jwtService)

	// Controllers
	statusController := controllers.NewStatusController()
	app.Get("/status", statusController.GetStatus)

	authController := controllers.NewAuthController(userRepository, jwtService)
	app.Post("/auth/sign-up", authController.SignUp)
	app.Post("/auth/sign-in", authController.SignIn)

	gistsController := controllers.NewGistsController(gistRepository)
	app.Get("/gists", authMiddleware.Handle, gistsController.GetGists)
	app.Post("/gists", authMiddleware.Handle, gistsController.CreateGist)
	app.Get("/gists/:id", authMiddleware.Handle, gistsController.GetGist)
	app.Put("/gists/:id", authMiddleware.Handle, gistsController.UpdateGist)
	app.Delete("/gists/:id", authMiddleware.Handle, gistsController.DeleteGist)

	filesController := controllers.NewFilesController(fileRepository, gistRepository)
	app.Get("/gists/:id/files", authMiddleware.Handle, filesController.GetFiles)
	app.Post("/gists/:id/files", authMiddleware.Handle, filesController.CreateFile)
	app.Get("/gists/:id/files/:fileId", authMiddleware.Handle, filesController.GetFile)
	app.Put("/gists/:id/files/:fileId", authMiddleware.Handle, filesController.UpdateFile)
	app.Delete("/gists/:id/files/:fileId", authMiddleware.Handle, filesController.DeleteFile)
}
