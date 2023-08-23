package routes

import (
	enrollmentHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/enrollment"
	"github.com/gofiber/fiber/v2"
)

func SetupEnrollmentRoutes(router fiber.Router) {
	media := router.Group("/enrollment")
	media.Post("/", enrollmentHandler.CreateEnrollment)
	media.Get("/", enrollmentHandler.GetEnrollments)
	media.Get("/:enrollmentId", enrollmentHandler.GetEnrollments)
	media.Put("/:enrollmentId", enrollmentHandler.UpdateEnrollmentByID)
	media.Delete("/:enrollmentId", enrollmentHandler.DeleteEnrollmentByID)
}
