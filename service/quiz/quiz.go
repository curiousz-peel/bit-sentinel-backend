package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetQuizzes() ([]models.QuizDTO, error) {
	quizzes := []models.Quiz{}
	quizzDTOs := []models.QuizDTO{}
	err := storage.DB.Model(&models.Quiz{}).Preload("Course").Preload("Lesson").Find(&quizzes).Error
	if err != nil {
		return nil, err
	}
	for i := range quizzes {
		if quizzes[i].QuestionIDs != nil {
			var ids []uint
			err = json.Unmarshal(quizzes[i].QuestionIDs, &ids)
			if err != nil {
				return nil, err
			}
			if len(ids) > 0 {
				res := storage.DB.Find(&quizzes[i].Questions, "id IN (?)", ids)
				if res.Error != nil || res.RowsAffected == 0 {
					return nil, errors.New("can't load questions for quiz with id " + fmt.Sprint(quizzes[i].ID))
				}
				storage.DB.Model(&quizzes[i]).Updates(&models.Quiz{
					Questions: quizzes[i].Questions})
			}
		}
	}
	for _, quiz := range quizzes {
		quizzDTOs = append(quizzDTOs, models.ToQuizDTO(quiz))
	}
	return quizzDTOs, nil
}

func GetQuizByID(id string) (*models.QuizDTO, error) {
	quiz := &models.Quiz{}
	res := storage.DB.Where("id = ?", id).Find(quiz)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("can't get quiz with id: " + id)
	}
	if quiz.QuestionIDs != nil {
		var ids []uint
		err := json.Unmarshal(quiz.QuestionIDs, &ids)
		if err != nil {
			return nil, err
		}
		res = storage.DB.Find(&quiz.Questions, "id IN (?)", ids)
		if res.Error != nil || res.RowsAffected == 0 {
			return nil, errors.New("could not find questions for quiz " + fmt.Sprint(quiz.ID))
		}
		storage.DB.Model(&quiz).Updates(&models.Quiz{
			Questions: quiz.Questions})
	}
	quizDTO := models.ToQuizDTO(*quiz)
	return &quizDTO, nil
}

func DeleteQuizByID(id string) error {
	quiz := &models.Quiz{}
	res := storage.DB.Unscoped().Delete(quiz, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete quiz with id " + id)
	}
	return nil
}

func UpdateQuizByID(id string, updateQuiz models.UpdateQuiz) error {
	quiz := &models.Quiz{}
	res := storage.DB.Where("id = ?", id).Find(quiz)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the quiz, check if ID " + id + " exists")
	}

	var questionIDsJSON []byte
	if updateQuiz.QuestionIDs != nil {
		questionIDsJSON, _ = json.Marshal(updateQuiz.QuestionIDs)
		res := storage.DB.Debug().Find(&quiz.Questions, "id IN (?)", updateQuiz.QuestionIDs)
		if res.Error != nil || res.RowsAffected == 0 || len(quiz.Questions) != len(updateQuiz.QuestionIDs) {
			return errors.New("error in finding questions when updating quiz: " +
				fmt.Sprint(len(quiz.Questions)) + " questions found, but " +
				fmt.Sprint(len(updateQuiz.QuestionIDs)) + " question IDs given")
		}
	}
	storage.DB.Model(&quiz).Updates(&models.Quiz{
		Title:       updateQuiz.Title,
		Description: updateQuiz.Description,
		CourseID:    updateQuiz.CourseID,
		LessonID:    updateQuiz.LessonID})

	if questionIDsJSON != nil {
		storage.DB.Model(&quiz).Updates(&models.Quiz{
			QuestionIDs: questionIDsJSON,
			Questions:   quiz.Questions})
	}
	return nil
}

func CreateQuiz(quiz *models.Quiz) (*models.QuizDTO, error) {
	err := storage.DB.Create(quiz).Error
	if err != nil {
		return nil, err
	}
	quizDTO := models.ToQuizDTO(*quiz)
	return &quizDTO, nil
}

func AddQuestionsToQuiz(id string, addedQuestionsIDs models.AddQuestionsToQuiz) error {
	quiz := &models.Quiz{}
	res := storage.DB.Where("id = ?", id).Find(quiz)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the quiz, check if ID " + id + " exists")
	}

	if addedQuestionsIDs.QuestionIDs == nil {
		return errors.New("no new questions to append to quizz " + id)
	}

	var existingQuestions []uint
	err := json.Unmarshal(quiz.QuestionIDs, &existingQuestions)
	if err != nil {
		return errors.New("can't unmarshal existing questions ids for quiz " + id)
	}
	var addedQuestions []models.Question
	res = storage.DB.Debug().Find(&addedQuestions, "id IN (?)", addedQuestionsIDs.QuestionIDs)
	if res.Error != nil || res.RowsAffected == 0 || len(addedQuestions) != len(addedQuestionsIDs.QuestionIDs) {
		return errors.New("error in finding options when appending to question: " +
			fmt.Sprint(len(addedQuestions)) + " questions found, but " +
			fmt.Sprint(len(addedQuestionsIDs.QuestionIDs)) + " question IDs given")
	}

	for _, questionId := range addedQuestionsIDs.QuestionIDs {
		idx := slices.Index(existingQuestions, questionId)
		if idx != -1 {
			return errors.New("can't have question with id " + fmt.Sprint(questionId) + " more than one time for quiz " + id)
		}
		existingQuestions = append(existingQuestions, questionId)
	}

	questionsToJSON, _ := json.Marshal(existingQuestions)
	storage.DB.Model(&quiz).Updates(&models.Quiz{
		QuestionIDs: questionsToJSON})
	return nil
}
