package service

import (
	"errors"
	"fmt"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetMedias() (*[]models.Media, error) {
	medias := &[]models.Media{}
	err := storage.DB.Model(&models.Media{}).Preload("Lesson").Preload("FileType").Find(&medias).Error
	if err != nil {
		return nil, err
	}
	return medias, nil
}

func GetMediaByID(id string) (*models.Media, error) {
	media := &models.Media{}
	res := storage.DB.Where("id = ?", id).Preload("Lesson").Preload("FileType").Find(media)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding media")
	}
	return media, nil
}

func DeleteMediaByID(id string) error {
	media := &models.Media{}
	res := storage.DB.Unscoped().Delete(media, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete media with id " + id)
	}
	return nil
}

func CreateMedia(media *models.Media) (*models.Media, error) {
	res := storage.DB.Find(media.Lesson, "id = ?", media.LessonID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding lesson with id: " + fmt.Sprint(media.LessonID))
	}
	res = storage.DB.Find(&media.FileType, "type = ?", media.FileTypeName)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding media format: " + media.FileTypeName)
	}
	err := storage.DB.Create(media).Error
	if err != nil {
		return nil, err
	}
	return media, nil
}

func UpdateMediaByID(id string, updateMedia models.UpdateMedia) error {
	media := &models.Media{}
	res := storage.DB.Where("id = ?", id).Preload("Lesson").Preload("FileType").Find(media)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not update media with id " + id)
	}
	storage.DB.Model(&media).Updates(&models.Media{
		LessonID:     updateMedia.LessonID,
		FilePath:     updateMedia.FilePath,
		FileTypeName: updateMedia.FileTypeName})
	return nil
}
