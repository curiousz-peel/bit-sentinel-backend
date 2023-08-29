package service

import (
	"errors"
	"fmt"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetEnrollments() (*[]models.EnrollmentDTO, error) {
	enrollments := &[]models.Enrollment{}
	enrollmentDTOs := []models.EnrollmentDTO{}

	err := storage.DB.Model(&models.Enrollment{}).Preload("User").Preload("Course").Preload("Progress").Find(&enrollments).Error
	if err != nil {
		return nil, err
	}
	for _, enrollment := range *enrollments {
		enrollmentDTOs = append(enrollmentDTOs, models.ToEnrollmentDTO(enrollment))
	}
	return &enrollmentDTOs, nil
}

func GetEnrollmentByID(id string) (*models.EnrollmentDTO, error) {
	enrollment := &models.Enrollment{}
	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Course").Preload("Progress").Find(enrollment)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("can't find enrollment with id " + id)
	}
	enrollmentDTO := models.ToEnrollmentDTO(*enrollment)
	return &enrollmentDTO, nil
}

func DeleteEnrollmentByID(id string) error {
	enrollment := &models.Enrollment{}
	res := storage.DB.Unscoped().Delete(enrollment, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete enrollment with id " + id)
	}
	return nil
}

func CreateEnrollment(enrollment *models.Enrollment) (*models.EnrollmentDTO, error) {
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
	enrollmentDTO := models.ToEnrollmentDTO(*enrollment)
	return &enrollmentDTO, nil
}

func UpdateEnrollmentByID(id string, updateEnrollment models.UpdateEnrollment) error {
	enrollment := &models.Enrollment{}
	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Course").Preload("Progress").Find(enrollment)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not update enrollment with id " + id)
	}
	storage.DB.Model(&enrollment).Updates(&models.Enrollment{
		ProgressID: uint(updateEnrollment.ProgressID),
		CourseID:   uint(updateEnrollment.CourseID),
		UserID:     updateEnrollment.UserID})
	return nil
}
