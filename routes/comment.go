package routes

import (
	commentHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/comment"
	"github.com/gofiber/fiber/v2"
)

func SetupCommentRoutes(router fiber.Router) {
	comment := router.Group("/comment")
	comment.Post("/", commentHandler.CreateComment)
	comment.Get("/", commentHandler.GetComments)
	comment.Get("/:commentId", commentHandler.GetCommentByID)
	comment.Put("/:commentId", commentHandler.UpdateCommentByID)
	comment.Delete("/:commentId", commentHandler.DeleteCommentByID)
}
