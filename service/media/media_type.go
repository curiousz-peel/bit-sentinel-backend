package service

import (
	"errors"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetMediaTypes() (*[]models.MediaType, error) {
	mediaTypes := &[]models.MediaType{}
	err := storage.DB.Model(&models.MediaType{}).Find(&mediaTypes).Error
	if err != nil {
		return nil, err
	}
	return mediaTypes, nil
}

func GetMediaTypeByID(id string) (*models.MediaType, error) {
	mediaType := &models.MediaType{}
	res := storage.DB.Where("id = ?", id).Find(mediaType)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding media type")
	}
	return mediaType, nil
}

func GetMediaTypeByType(mediaTypeName string) (*models.MediaType, error) {
	mediaType := &models.MediaType{}
	res := storage.DB.Where("type = ?", mediaTypeName).Find(mediaType)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding media type")
	}
	return mediaType, nil
}

func CreateMediaType(mediaType *models.MediaType) (*models.MediaType, error) {
	err := storage.DB.Create(mediaType).Error
	if err != nil {
		return nil, err
	}
	return mediaType, nil
}

func DeleteMediaTypeByID(id string) error {
	media := &models.MediaType{}
	res := storage.DB.Unscoped().Delete(media, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("error in deleting media type")
	}
	return nil
}

func UpdateMediaTypeByID(id string, updateMediaType models.UpdateMediaType) error {
	mediaType := &models.MediaType{}
	res := storage.DB.Where("id = ?", id).Find(mediaType)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("can't find media type to update")
	}
	storage.DB.Model(&mediaType).Updates(&models.MediaType{
		Type: updateMediaType.Type})
	return nil
}
