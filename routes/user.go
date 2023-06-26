package routes

import (
	userHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/user"
	"github.com/gofiber/fiber/v2"
)

// type User struct {
// 	ID        int       `json:"id"`
// 	FirstName string    `json:"firstName"`
// 	LastName  string    `json:"lastName"`
// 	Email     string    `json:"email"`
// 	Password  string    `json:"password"`
// 	Birthday  time.Time `json:"birthday"`
// 	IsMod     bool      `json:"isModerator"`
// }

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")
	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.GetUsers)
	user.Get("/:userId", userHandler.GetUserByID)
	user.Put("/:userId", userHandler.UpdateUserByID)
	user.Delete("/:userId", userHandler.DeleteUserByID)
}
