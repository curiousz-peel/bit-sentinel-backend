package routes

import (
	commentHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/comment"
	"github.com/gofiber/fiber/v2"
)

func SetupCommentRoutes(router fiber.Router) {
	media := router.Group("/comment")
	media.Post("/", commentHandler.CreateComment)
	media.Get("/", commentHandler.GetComments)
	media.Get("/:commentId", commentHandler.GetCommentByID)
	media.Put("/:commentId", commentHandler.UpdateCommentByID)
	media.Delete("/:commentId", commentHandler.DeleteCommentByID)
}
