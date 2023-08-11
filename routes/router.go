package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	SetupUserRoutes(api)
	SetupSubscriptionRoutes(api)
	SetupAuthorRoutes(api)
	SetupSubscriptionPlanRoutes(api)
	// SetupMediaRoutes(api)
}
