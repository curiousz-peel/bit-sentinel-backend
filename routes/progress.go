package routes

import (
	progressHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/progress"
	"github.com/gofiber/fiber/v2"
)

func SetupProgressRoutes(router fiber.Router) {
	progress := router.Group("/progress")
	progress.Post("/", progressHandler.CreateProgress)
	progress.Get("/", progressHandler.GetProgresss)
	progress.Get("/:progressId", progressHandler.GetProgressByID)
	progress.Put("/:progressId", progressHandler.UpdateProgressByID)
	progress.Delete("/:progressId", progressHandler.DeleteProgressByID)
}
