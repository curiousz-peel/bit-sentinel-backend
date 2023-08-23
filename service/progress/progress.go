package service

import (
	"errors"
	"fmt"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetProgresss() (*[]models.Progress, error) {
	progresss := &[]models.Progress{}
	err := storage.DB.Model(&models.Progress{}).Find(&progresss).Error
	if err != nil {
		return nil, err
	}
	return progresss, nil
}

func GetProgressByID(id string) (*models.Progress, error) {
	progress := &models.Progress{}
	res := storage.DB.Where("id = ?", id).Find(progress)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("could not find progress with id " + id)
	}
	return progress, nil
}

func DeleteProgressByID(id string) error {
	progress := &models.Progress{}
	res := storage.DB.Unscoped().Delete(progress, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete progress with id " + id)
	}
	return nil
}

func CreateProgress(progress *models.Progress) (*models.Progress, error) {
	enrollment := models.Enrollment{}
	res := storage.DB.Debug().Find(enrollment, "id = ?", progress.EnrollmentID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding enrollment with id: " + fmt.Sprint(progress.EnrollmentID))
	}

	err := storage.DB.Debug().Create(progress).Error
	if err != nil {
		return nil, err
	}
	return progress, nil
}

func UpdateProgressByID(id string, updateProgress models.UpdateProgress) error {
	progress := &models.Progress{}
	res := storage.DB.Where("id = ?", id).Find(progress)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the progress with id " + id)
	}
	storage.DB.Model(&progress).Updates(&models.Progress{
		Completed: updateProgress.Completed,
		Progress:  updateProgress.Progress})
	return nil
}
