package routes

import (
	authorHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/author"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthorRoutes(router fiber.Router) {
	author := router.Group("/author")
	author.Post("/", authorHandler.CreateAuthor)
	author.Get("/", jwtHandler.ValidateToken, authorHandler.GetAuthors)
	author.Get("/:authorId", authorHandler.GetAuthorByID)
	author.Delete("/:authorId", authorHandler.DeleteAuthorByID)
	author.Put("/:authorId", authorHandler.UpdateAuthorByID)
}
