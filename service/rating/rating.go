package service

import (
	"errors"
	"fmt"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetRatings() (*[]models.RatingDTO, error) {
	ratings := &[]models.Rating{}
	ratingsDTOs := []models.RatingDTO{}
	err := storage.DB.Model(&models.Rating{}).Preload("User").Preload("Course").Find(&ratings).Error
	if err != nil {
		return nil, err
	}
	for _, rating := range *ratings {
		ratingsDTOs = append(ratingsDTOs, models.ToRatingDTO(rating))
	}
	return &ratingsDTOs, nil
}

func GetRatingByID(id string) (*models.RatingDTO, error) {
	rating := &models.Rating{}
	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Course").Find(rating)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("could not find rating with id " + id)
	}
	ratingDTO := models.ToRatingDTO(*rating)
	return &ratingDTO, nil
}

func DeleteRatingByID(id string) error {
	rating := &models.Rating{}
	res := storage.DB.Unscoped().Delete(rating, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete rating with id " + id)
	}
	return nil
}

func CreateRating(rating *models.Rating) (*models.RatingDTO, error) {
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
	ratingDTO := models.ToRatingDTO(*rating)
	return &ratingDTO, nil
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
