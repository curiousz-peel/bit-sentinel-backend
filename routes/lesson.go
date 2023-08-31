package routes

import (
	lessonHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/lesson"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func SetupLessonRoutes(router fiber.Router) {
	lesson := router.Group("/lesson")
	lesson.Post("/", jwtHandler.ValidateToken, lessonHandler.CreateLesson)
	lesson.Get("/", jwtHandler.ValidateToken, lessonHandler.GetLessons)
	lesson.Get("/:lessonId", jwtHandler.ValidateToken, lessonHandler.GetLessonByID)
	lesson.Put("/:lessonId", jwtHandler.ValidateToken, lessonHandler.UpdateLessonByID)
	lesson.Delete("/:lessonId", jwtHandler.ValidateToken, lessonHandler.DeleteLessonByID)
}
