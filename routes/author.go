package routes

import (
	authorHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/author"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthorRoutes(router fiber.Router) {
	author := router.Group("/author")
	author.Post("/", jwtHandler.ValidateToken, authorHandler.CreateAuthor)
	author.Get("/", jwtHandler.ValidateToken, authorHandler.GetAuthors)
	author.Get("/:authorId", jwtHandler.ValidateToken, authorHandler.GetAuthorByID)
	author.Delete("/:authorId", jwtHandler.ValidateToken, authorHandler.DeleteAuthorByID)
	author.Put("/:authorId", jwtHandler.ValidateToken, authorHandler.UpdateAuthorByID)
}
