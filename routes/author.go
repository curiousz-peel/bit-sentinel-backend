package routes

import (
	userHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/author"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthorRoutes(router fiber.Router) {
	user := router.Group("/author")
	user.Post("/", userHandler.CreateAuthor)
	user.Get("/", userHandler.GetAuthors)
}
