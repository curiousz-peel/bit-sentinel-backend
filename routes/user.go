package routes

import (
	userHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/user"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")
	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.GetUsers)
	user.Get("/:userId", userHandler.GetUserByID)
	user.Put("/:userId", userHandler.UpdateUserByID)
	user.Delete("/:userId", userHandler.DeleteUserByID)
}
