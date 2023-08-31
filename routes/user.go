package routes

import (
	userHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/user"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")
	user.Post("/", jwtHandler.ValidateToken, userHandler.CreateUser)
	user.Get("/", jwtHandler.ValidateToken, userHandler.GetUsers)
	user.Get("/:userId", jwtHandler.ValidateToken, userHandler.GetUserByID)
	user.Put("/:userId", jwtHandler.ValidateToken, userHandler.UpdateUserByID)
	user.Delete("/:userId", jwtHandler.ValidateToken, userHandler.DeleteUserByID)
}
