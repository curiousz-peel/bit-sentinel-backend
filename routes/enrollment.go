package routes

import (
	enrollmentHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/enrollment"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func SetupEnrollmentRoutes(router fiber.Router) {
	media := router.Group("/enrollment")
	media.Post("/", jwtHandler.ValidateToken, enrollmentHandler.CreateEnrollment)
	media.Get("/", jwtHandler.ValidateToken, enrollmentHandler.GetEnrollments)
	media.Get("/:enrollmentId", jwtHandler.ValidateToken, enrollmentHandler.GetEnrollments)
	media.Put("/:enrollmentId", jwtHandler.ValidateToken, enrollmentHandler.UpdateEnrollmentByID)
	media.Delete("/:enrollmentId", jwtHandler.ValidateToken, enrollmentHandler.DeleteEnrollmentByID)
}
