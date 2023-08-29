package service

import (
	"encoding/json"
	"errors"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetAuthorByID(id string) (*models.AuthorDTO, error) {
	author := &models.Author{}
	res := storage.DB.Where("id = ?", id).Preload("User").Find(author)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("could not find author with id " + id)
	}
	authorDTO := models.ToAuthorDTO(*author)
	return &authorDTO, nil
}

func DeleteAuthorByID(id string) error {
	author := &models.Author{}
	res := storage.DB.Where("id = ?", id).Unscoped().Delete(&author)
	// res := storage.DB.Unscoped().Delete(author, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete author with id " + id)
	}
	return nil
}

func CreateAuthor(author *models.Author) (*models.AuthorDTO, error) {
	res := storage.DB.Debug().Find(&author.User, "id = ?", author.UserID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding user: " + author.UserID.String())
	}
	err := storage.DB.Debug().Omit("User").Create(author).Error
	if err != nil {
		return nil, err
	}
	res = storage.DB.Model(author.User).Update("is_author", true)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("can't update author status for user with id " + author.UserID.String())
	}

	authorDTO := models.ToAuthorDTO(*author)
	return &authorDTO, nil
}

func UpdateAuthorByID(id string, updateAuthor models.UpdateAuthor) error {
	author := &models.Author{}
	res := storage.DB.Where("id = ?", id).Preload("User").Find(author)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the author with id " + id)
	}
	topicsJSON, err := json.Marshal(updateAuthor.Topics)
	if err != nil {
		return errors.New("can't marshal topics with err " + err.Error())
	}
	storage.DB.Model(&author).Updates(&models.Author{
		Profession:  updateAuthor.Profession,
		Description: updateAuthor.Description,
		Topics:      topicsJSON,
		UserID:      updateAuthor.UserID})
	return nil
}

func GetAuthors() (*[]models.AuthorDTO, error) {
	authors := &[]models.Author{}
	authorDTOs := []models.AuthorDTO{}
	err := storage.DB.Model(&models.Author{}).Preload("User").Find(&authors).Error
	if err != nil {
		return nil, err
	}
	for _, author := range *authors {
		authorDTOs = append(authorDTOs, models.ToAuthorDTO(author))
	}
	return &authorDTOs, nil
}
