package routes

import (
	ratingHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/rating"
	"github.com/gofiber/fiber/v2"
)

func SetupRatingRoutes(router fiber.Router) {
	rating := router.Group("/rating")
	rating.Post("/", ratingHandler.CreateRating)
	rating.Get("/", ratingHandler.GetRatings)
	rating.Get("/:ratingId", ratingHandler.GetRatingByID)
	rating.Put("/:ratingId", ratingHandler.UpdateRatingByID)
	rating.Delete("/:ratingId", ratingHandler.DeleteRatingByID)
}
