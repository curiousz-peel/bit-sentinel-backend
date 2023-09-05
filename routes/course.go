package routes

import (
	courseHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/course"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func SetupCourseRoutes(router fiber.Router) {
	course := router.Group("/course")
	course.Post("/", jwtHandler.ValidateToken, courseHandler.CreateCourse)
	course.Get("/", jwtHandler.ValidateToken, courseHandler.GetCourses)
	course.Get("/recent", jwtHandler.ValidateToken, courseHandler.GetCoursesByMostRecentForHome)
	course.Get("/rating", jwtHandler.ValidateToken, courseHandler.GetCoursesByRatingForHome)
	course.Get("/fundamental", jwtHandler.ValidateToken, courseHandler.GetCoursesFundamentalsForHome)
	course.Get("/tag/:value", jwtHandler.ValidateToken, courseHandler.GetCoursesByTag)
	course.Get("/:courseId", jwtHandler.ValidateToken, courseHandler.GetCourseByID)
	course.Get("/subscription/:subscriptionType", jwtHandler.ValidateToken, courseHandler.GetCoursesBySubscription)
	course.Get("/author/:authorId", jwtHandler.ValidateToken, courseHandler.GetCoursesByAuthorId)
	course.Put("/:courseId", jwtHandler.ValidateToken, courseHandler.UpdateCourseByID)
	course.Put("/addAuthors:courseId", jwtHandler.ValidateToken, courseHandler.AddAuthorsToCourse)
	course.Delete("/:courseId", jwtHandler.ValidateToken, courseHandler.DeleteCourseByID)
}
