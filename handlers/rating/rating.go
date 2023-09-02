package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/rating"
	"github.com/gofiber/fiber/v2"
)

func GetRatings(ctx *fiber.Ctx) error {
	ratings, err := service.GetRatings()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch ratings",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "ratings fetched successfully",
		"data":    ratings})

}

func GetRatingByID(ctx *fiber.Ctx) error {
	id := ctx.Params("ratingId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "rating ID cannot be empty on get",
			"data":    nil})
	}
	rating, err := service.GetRatingByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the rating, check if rating with ID " + id + " exists",
			"data":    err.Error()})
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "rating fetched successfully",
		"data":    rating})
	return nil
}

func GetRatingsByCourseID(ctx *fiber.Ctx) error {
	id := ctx.Params("courseId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "course ID cannot be empty on get ratings by course ID",
			"data":    nil})
	}
	rating, err := service.GetRatingsByCourseId(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "errors in fetching the ratings",
			"data":    err.Error()})
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "ratings fetched successfully",
		"data":    rating})
	return nil
}

func DeleteRatingByID(ctx *fiber.Ctx) error {
	id := ctx.Params("ratingId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "rating ID cannot be empty on delete",
			"data":    nil})
	}
	err := service.DeleteRatingByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete rating, check if rating with ID " + id + " exists",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "rating deleted successfully"})
}

func CreateRating(ctx *fiber.Ctx) error {
	rating := &models.Rating{}
	err := ctx.BodyParser(rating)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err})
	}
	ratingDTO, err := service.CreateRating(rating)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create rating",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "rating created successfully",
		"data":    ratingDTO})
}

func UpdateRatingByID(ctx *fiber.Ctx) error {
	id := ctx.Params("ratingId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "rating ID cannot be empty on update",
			"data":    nil})
	}
	var updateRatingData models.UpdateRating
	err := ctx.BodyParser(&updateRatingData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the user id, the course id or the rating value",
			"data":    err})
	}
	err = service.UpdateRatingByID(id, updateRatingData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the rating",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "rating updated successfully",
		"data":    nil})
}
