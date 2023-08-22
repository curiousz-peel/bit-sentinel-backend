package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strconv"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetOptions() (*[]models.Option, error) {
	options := &[]models.Option{}
	err := storage.DB.Model(&models.Option{}).Preload("Question").Find(&options).Error
	if err != nil {
		return nil, err
	}
	return options, nil
}

func GetOptionByID(id string) (*models.Option, error) {
	option := &models.Option{}
	res := storage.DB.Where("id = ?", id).Preload("Question.Quiz").Find(option)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("couldn't find option with id " + id)
	}
	return option, nil
}

func DeleteOptionByID(id string) error {
	option := &models.Option{}

	questions := []models.Question{}
	intId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("can't parse id " + id + " to int with err: " + err.Error())
	}

	query := fmt.Sprintf("SELECT * FROM questions WHERE option_ids @> '[%d]'", intId)
	err = storage.DB.Raw(query).Scan(&questions).Error
	if err != nil {
		return errors.New("can't query options from questions with err: " + err.Error())
	}

	for _, question := range questions {
		var ids []uint
		err = json.Unmarshal(question.OptionIDs, &ids)
		if err != nil {
			return errors.New("unmarshaling ids failed with error: " + err.Error())
		}

		idx := slices.Index(ids, uint(intId))
		newOptionIDs := slices.Delete(ids, idx, idx+1)

		if len(newOptionIDs) == 0 {
			DeleteQuestionByID(strconv.FormatUint(uint64(question.ID), 10))
		} else {
			updateQuestionOptionIDs := models.UpdateQuestion{OptionIDs: newOptionIDs}
			err = UpdateQuestionByID(strconv.FormatUint(uint64(question.ID), 10), updateQuestionOptionIDs)
			if err != nil {
				return err
			}
		}
	}
	res := storage.DB.Unscoped().Delete(option, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return res.Error
	}
	return nil
}

func CreateOption(option *models.Option) (*models.DisplayOption, error) {
	res := storage.DB.Find(&option.Question, "id = ?", option.QuestionID)
	if res.Error != nil {
		return nil, errors.New("error in finding question: " + fmt.Sprint(option.QuestionID))
	}

	err := storage.DB.Create(option).Error
	if err != nil {
		return nil, err
	}
	return &models.DisplayOption{ID: option.ID, QuestionID: option.QuestionID, IsCorrect: option.IsCorrect, Text: option.Text}, nil
}

func UpdateOptionByID(id string, updateOption models.UpdateOption) error {
	option := &models.Option{}
	res := storage.DB.Where("id = ?", id).Preload("Question").Find(option)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the option with id " + id)
	}

	storage.DB.Model(&option).Updates(&models.Option{
		Text:       updateOption.Text,
		QuestionID: updateOption.QuestionID,
		IsCorrect:  updateOption.IsCorrect})
	return nil
}
