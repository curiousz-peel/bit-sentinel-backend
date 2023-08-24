package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	SetupUserRoutes(api)
	SetupSubscriptionRelatedRoutes(api)
	SetupAuthorRoutes(api)
	SetupTestComplexRoutes(api)
	SetupMediaRelatedRoutes(api)
	SetupCommentRoutes(api)
	SetupEnrollmentRoutes(api)
	SetupQuizRelatedRoutes(api)
	SetupRatingRoutes(api)
	SetupProgressRoutes(api)
	SetupLessonRoutes(api)
}
