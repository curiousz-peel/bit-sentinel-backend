package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api") // Group endpoints with param 'api' and log whenever this endpoint is hit.
	SetupUserRoutes(api)
}
