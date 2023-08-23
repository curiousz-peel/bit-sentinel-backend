package service

import (
	"errors"
	"fmt"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetEnrollments() (*[]models.Enrollment, error) {
	enrollments := &[]models.Enrollment{}
	err := storage.DB.Model(&models.Enrollment{}).Preload("User").Preload("Course").Preload("Progress").Find(&enrollments).Error
	if err != nil {
		return nil, err
	}
	return enrollments, nil
}

func GetEnrollmentByID(id string) (*models.Enrollment, error) {
	enrollment := &models.Enrollment{}
	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Course").Preload("Progress").Find(enrollment)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, res.Error
	}
	return enrollment, nil
}

func DeleteEnrollmentByID(id string) error {
	enrollment := &models.Enrollment{}
	res := storage.DB.Unscoped().Delete(enrollment, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return res.Error
	}
	return nil
}

func CreateEnrollment(enrollment *models.Enrollment) (*models.Enrollment, error) {
	res := storage.DB.Find(&enrollment.User, "id = ?", enrollment.UserID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding user: " + enrollment.UserID.String())
	}

	res = storage.DB.Find(&enrollment.Course, "id = ?", enrollment.CourseID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding course: " + fmt.Sprint(enrollment.CourseID))
	}

	res = storage.DB.Find(&enrollment.Progress, "id = ?", enrollment.ProgressID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding progress: " + fmt.Sprint(enrollment.ProgressID))
	}

	err := storage.DB.Omit("User").Create(enrollment).Error
	if err != nil {
		return nil, err
	}
	return enrollment, nil
}

func UpdateEnrollmentByID(id string, updateEnrollment models.UpdateEnrollment) error {
	enrollment := &models.Enrollment{}
	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Course").Preload("Progress").Find(enrollment)
	if res.Error != nil || res.RowsAffected == 0 {
		return res.Error
	}
	storage.DB.Model(&enrollment).Updates(&models.Enrollment{
		ProgressID: uint(updateEnrollment.ProgressID),
		CourseID:   uint(updateEnrollment.CourseID),
		UserID:     updateEnrollment.UserID})
	return nil
}
