package service

import (
	"errors"
	"fmt"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetRatings() (*[]models.Rating, error) {
	ratings := &[]models.Rating{}
	err := storage.DB.Model(&models.Rating{}).Preload("User").Preload("Course").Find(&ratings).Error
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func GetRatingByID(id string) (*models.Rating, error) {
	rating := &models.Rating{}
	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Course").Find(rating)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("could not find rating with id " + id)
	}
	return rating, nil
}

func DeleteRatingByID(id string) error {
	rating := &models.Rating{}
	res := storage.DB.Unscoped().Delete(rating, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete rating with id " + id)
	}
	return nil
}

func CreateRating(rating *models.Rating) (*models.Rating, error) {
	res := storage.DB.Debug().Find(&rating.User, "id = ?", rating.UserID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding user: " + rating.UserID.String())
	}

	if rating.CourseID != 0 {
		res = storage.DB.Find(&rating.Course, "id = ?", rating.CourseID)
		if res.Error != nil || res.RowsAffected == 0 {
			return nil, errors.New("error in finding course: " + fmt.Sprint(rating.CourseID))
		}
	}

	err := storage.DB.Debug().Omit("User").Create(rating).Error
	if err != nil {
		return nil, err
	}
	return rating, nil
}

func UpdateRatingByID(id string, updateRating models.UpdateRating) error {
	rating := &models.Rating{}
	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Course").Find(rating)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the rating with id " + id)
	}
	storage.DB.Model(&rating).Updates(&models.Rating{
		Rating:   updateRating.Rating,
		CourseID: uint(updateRating.CourseID),
		UserID:   updateRating.UserID})

	return nil
}

// func UpdateRatingByID(ctx *fiber.Ctx) error {

// 	rating := &models.Rating{}
// 	id := ctx.Params("ratingId")
// 	if id == "" {
// 		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
// 			"message": "rating ID cannot be empty on update",
// 			"data":    nil})
// 	}

// 	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Course").Find(rating)
// 	if res.Error != nil || res.RowsAffected == 0 {
// 		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
// 			"message": "could not find the rating",
// 			"data":    res.Error})
// 		return res.Error
// 	}

// 	var updateRatingData updateRating
// 	err := ctx.BodyParser(&updateRatingData)
// 	if err != nil {
// 		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"message": "bad input: you can only update the user id, the course id or the rating value",
// 			"data":    err})
// 		return err
// 	}

// 	type UpdateRatingResponse struct {
// 		CourseID uint
// 		UserID   string
// 		Rating   float64
// 	}

// 	ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "rating updated successfully",
// 		"data":    UpdateRatingResponse{CourseID: rating.CourseID, UserID: rating.UserID.String(), Rating: rating.Rating}})
// 	return nil
// }
