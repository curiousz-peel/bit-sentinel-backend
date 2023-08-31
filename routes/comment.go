package routes

import (
	commentHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/comment"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func SetupCommentRoutes(router fiber.Router) {
	comment := router.Group("/comment")
	comment.Post("/", jwtHandler.ValidateToken, commentHandler.CreateComment)
	comment.Get("/", jwtHandler.ValidateToken, commentHandler.GetComments)
	comment.Get("/:commentId", jwtHandler.ValidateToken, commentHandler.GetCommentByID)
	comment.Put("/:commentId", jwtHandler.ValidateToken, commentHandler.UpdateCommentByID)
	comment.Delete("/:commentId", jwtHandler.ValidateToken, commentHandler.DeleteCommentByID)
}
