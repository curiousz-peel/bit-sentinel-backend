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

func GetQuestions() ([]models.QuestionDTO, error) {
	questions := []models.Question{}
	questionDTOs := []models.QuestionDTO{}
	err := storage.DB.Preload("Quiz").Find(&questions).Error
	if err != nil {
		return nil, err
	}
	for i := range questions {
		if questions[i].OptionIDs != nil {
			var ids []uint
			err = json.Unmarshal(questions[i].OptionIDs, &ids)
			if err != nil {
				return nil, err
			}
			if len(ids) > 0 {
				res := storage.DB.Find(&questions[i].Options, "id IN (?)", ids)
				if res.Error != nil || res.RowsAffected == 0 {
					return nil, errors.New("can't load options for question with id " + fmt.Sprint(questions[i].ID))
				}
				storage.DB.Model(&questions[i]).Updates(&models.Question{
					Options: questions[i].Options})
			}
		}
	}

	for _, question := range questions {
		questionDTOs = append(questionDTOs, models.ToQuestionDTO(question))
	}
	return questionDTOs, nil
}

func GetQuestionByID(id string) (*models.QuestionDTO, error) {
	question := &models.Question{}
	res := storage.DB.Preload("Quiz").Where("id = ?", id).Find(question)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("can't find question with id " + id)
	}
	if question.OptionIDs != nil {
		var ids []uint
		err := json.Unmarshal(question.OptionIDs, &ids)
		if err != nil {
			return nil, err
		}
		res = storage.DB.Find(&question.Options, "id IN (?)", ids)
		if res.Error != nil || res.RowsAffected == 0 {
			return nil, errors.New("could not find options for question " + fmt.Sprint(question.ID))
		}
	}
	storage.DB.Model(&question).Updates(&models.Question{
		Options: question.Options})

	questionDTO := models.ToQuestionDTO(*question)
	return &questionDTO, nil
}

func DeleteQuestionByID(id string) error {
	question := &models.Question{}

	quizzes := []models.Quiz{}
	intId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("can't parse id " + id + " to int with err: " + err.Error())
	}

	query := fmt.Sprintf("SELECT * FROM quizzes WHERE question_ids @> '[%d]'", intId)
	err = storage.DB.Raw(query).Scan(&quizzes).Error
	if err != nil {
		return errors.New("can't query questions from quizzes with err: " + err.Error())
	}

	for _, quiz := range quizzes {
		var ids []uint
		err = json.Unmarshal(quiz.QuestionIDs, &ids)
		if err != nil {
			return errors.New("unmarshaling ids failed with error: " + err.Error())
		}

		idx := slices.Index(ids, uint(intId))
		newQuestionIDs := slices.Delete(ids, idx, idx+1)

		if len(newQuestionIDs) == 0 {
			DeleteQuizByID(strconv.FormatUint(uint64(quiz.ID), 10))
		} else {
			updateQuizQuestionIDs := models.UpdateQuiz{QuestionIDs: newQuestionIDs}
			err = UpdateQuizByID(strconv.FormatUint(uint64(quiz.ID), 10), updateQuizQuestionIDs)
			if err != nil {
				return err
			}
		}
	}
	res := storage.DB.Unscoped().Delete(question, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete question with id " + id)
	}
	return nil
}

func UpdateQuestionByID(id string, updateQuestion models.UpdateQuestion) error {
	question := &models.Question{}
	res := storage.DB.Where("id = ?", id).Find(question)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the question, check if ID " + id + " exists")
	}

	var optionIDsJSON []byte
	if updateQuestion.OptionIDs != nil {
		optionIDsJSON, _ = json.Marshal(updateQuestion.OptionIDs)
		res := storage.DB.Debug().Find(&question.Options, "id IN (?)", updateQuestion.OptionIDs)
		if res.Error != nil || res.RowsAffected == 0 || len(question.Options) != len(updateQuestion.OptionIDs) {
			return errors.New("error in finding options when updating question: " +
				fmt.Sprint(len(question.Options)) + " questions found, but " +
				fmt.Sprint(len(updateQuestion.OptionIDs)) + " question IDs given")
		}
	}
	storage.DB.Model(&question).Updates(&models.Question{
		Text:   updateQuestion.Text,
		QuizID: updateQuestion.QuizID})

	if optionIDsJSON != nil {
		storage.DB.Model(&question).Updates(&models.Question{
			OptionIDs: optionIDsJSON,
			Options:   question.Options})
	}
	return nil
}

func AddOptionsToQuestion(id string, addedOptionsIDs models.AddOptionsToQuestion) error {
	question := &models.Question{}
	res := storage.DB.Where("id = ?", id).Find(question)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the question, check if ID " + id + " exists")
	}

	if addedOptionsIDs.OptionIDs == nil {
		return errors.New("no new options to append to question " + id)
	}

	var existingOptions []uint
	err := json.Unmarshal(question.OptionIDs, &existingOptions)
	if err != nil {
		return errors.New("can't unmarshal existing option ids for question " + id)
	}
	var addedOptions []models.Option
	res = storage.DB.Debug().Find(&addedOptions, "id IN (?)", addedOptionsIDs.OptionIDs)
	if res.Error != nil || res.RowsAffected == 0 || len(addedOptions) != len(addedOptionsIDs.OptionIDs) {
		return errors.New("error in finding options when appending to question: " +
			fmt.Sprint(len(addedOptions)) + " questions found, but " +
			fmt.Sprint(len(addedOptionsIDs.OptionIDs)) + " question IDs given")
	}

	for _, option := range addedOptionsIDs.OptionIDs {
		idx := slices.Index(existingOptions, option)
		if idx != -1 {
			return errors.New("can't have option with id " + fmt.Sprint(option) + " more than one time for question " + id)
		}
		existingOptions = append(existingOptions, option)
	}

	optionsToJSON, _ := json.Marshal(existingOptions)
	storage.DB.Model(&question).Updates(&models.Question{
		OptionIDs: optionsToJSON})
	return nil
}

func CreateQuestion(question *models.Question) (*models.QuestionDTO, error) {
	res := storage.DB.Model(models.Quiz{}).Find(&question.Quiz, "id = ?", question.QuizID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding quiz with id " + fmt.Sprint(question.QuizID) + " when creating question")
	}
	err := storage.DB.Create(question).Error
	if err != nil {
		return nil, err
	}
	questionDTO := models.ToQuestionDTO(*question)
	return &questionDTO, nil
}
