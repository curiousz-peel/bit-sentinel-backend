package routes

import (
	ratingHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/rating"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func SetupRatingRoutes(router fiber.Router) {
	rating := router.Group("/rating")
	rating.Post("/", jwtHandler.ValidateToken, ratingHandler.CreateRating)
	rating.Get("/", jwtHandler.ValidateToken, ratingHandler.GetRatings)
	rating.Get("/:ratingId", jwtHandler.ValidateToken, ratingHandler.GetRatingByID)
	rating.Get("/course/:courseId", jwtHandler.ValidateToken, ratingHandler.GetRatingsByCourseID)
	rating.Put("/:ratingId", jwtHandler.ValidateToken, ratingHandler.UpdateRatingByID)
	rating.Delete("/:ratingId", jwtHandler.ValidateToken, ratingHandler.DeleteRatingByID)
}
