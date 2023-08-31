package routes

import (
	progressHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/progress"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func SetupProgressRoutes(router fiber.Router) {
	progress := router.Group("/progress")
	progress.Post("/", jwtHandler.ValidateToken, progressHandler.CreateProgress)
	progress.Get("/", jwtHandler.ValidateToken, progressHandler.GetProgresss)
	progress.Get("/:progressId", jwtHandler.ValidateToken, progressHandler.GetProgressByID)
	progress.Put("/:progressId", jwtHandler.ValidateToken, progressHandler.UpdateProgressByID)
	progress.Delete("/:progressId", jwtHandler.ValidateToken, progressHandler.DeleteProgressByID)
}
