package routes

import (
	lessonHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/lesson"
	"github.com/gofiber/fiber/v2"
)

func SetupLessonRoutes(router fiber.Router) {
	lesson := router.Group("/lesson")
	lesson.Post("/", lessonHandler.CreateLesson)
	lesson.Get("/", lessonHandler.GetLessons)
	lesson.Get("/:lessonId", lessonHandler.GetLessonByID)
	lesson.Put("/:lessonId", lessonHandler.UpdateLessonByID)
	lesson.Delete("/:lessonId", lessonHandler.DeleteLessonByID)
}
