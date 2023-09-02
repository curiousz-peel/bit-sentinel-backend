package routes

import (
	loginHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/auth"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/login/", loginHandler.Login)
	auth.Post("/signup/", loginHandler.Signup)
}
