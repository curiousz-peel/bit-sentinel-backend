package service

import (
	"errors"
	"fmt"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetComments() (*[]models.CommentDTO, error) {
	comments := &[]models.Comment{}
	commentsDTO := []models.CommentDTO{}
	err := storage.DB.Model(&models.Comment{}).Preload("User").Preload("Course").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	for _, comment := range *comments {
		commentsDTO = append(commentsDTO, models.ToCommentDTO(comment))
	}
	return &commentsDTO, nil
}

func GetCommentByID(id string) (*models.CommentDTO, error) {
	comment := &models.Comment{}
	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Course").Find(comment)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("could not find comment with id " + id)
	}
	commentDTO := models.ToCommentDTO(*comment)
	return &commentDTO, nil
}

func DeleteCommentByID(id string) error {
	comment := &models.Comment{}
	res := storage.DB.Unscoped().Delete(comment, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete comment with id " + id)
	}
	return nil
}

func CreateComment(comment *models.Comment) (*models.CommentDTO, error) {
	res := storage.DB.Debug().Find(&comment.User, "id = ?", comment.UserID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding user: " + comment.UserID.String())
	}
	if comment.CourseID != 0 {
		res = storage.DB.Find(&comment.Course, "id = ?", comment.CourseID)
		if res.Error != nil || res.RowsAffected == 0 {
			return nil, errors.New("error in finding course: " + fmt.Sprint(comment.CourseID))
		}
	}
	err := storage.DB.Debug().Omit("User").Create(comment).Error
	if err != nil {
		return nil, err
	}
	commentDTO := models.ToCommentDTO(*comment)
	return &commentDTO, nil
}

func UpdateCommentByID(id string, updateComment models.UpdateComment) error {
	comment := &models.Comment{}
	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Course").Find(comment)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the comment with id " + id)
	}

	storage.DB.Model(&comment).Updates(&models.Comment{
		Text:     updateComment.Text,
		CourseID: uint(updateComment.CourseID),
		UserID:   updateComment.UserID})
	return nil
}
