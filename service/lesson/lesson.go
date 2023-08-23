package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strconv"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	courseService "github.com/curiousz-peel/web-learning-platform-backend/service/course"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetLessons() (*[]models.Lesson, error) {
	lessons := []models.Lesson{}
	err := storage.DB.Preload("Course").Find(&lessons).Error
	if err != nil {
		return nil, err
	}
	for _, lesson := range lessons {
		if lesson.ContentIds != nil {
			var ids []uint
			err = json.Unmarshal(lesson.ContentIds, &ids)
			if err != nil {
				return nil, err
			}
			if len(ids) > 0 {
				res := storage.DB.Find(&lesson.Content, "id IN (?)", ids)
				if res.Error != nil || res.RowsAffected == 0 {
					return nil, errors.New("can't load medias for lesson with id " + fmt.Sprint(lesson.ID))
				}
				storage.DB.Model(&lesson).Updates(&models.Lesson{
					Content: lesson.Content})
			}
		}
		storage.DB.Where("lesson_id = ?", lesson.ID).Find(&lesson.Quizzes)
	}
	return &lessons, nil
}

func GetLessonByID(id string) (*models.Lesson, error) {
	lesson := &models.Lesson{}
	res := storage.DB.Preload("Course").Where("id = ?", id).Find(lesson)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("can't find lesson with id " + id)
	}
	if lesson.ContentIds != nil {
		var ids []uint
		err := json.Unmarshal(lesson.ContentIds, &ids)
		if err != nil {
			return nil, err
		}
		res = storage.DB.Find(&lesson.Content, "id IN (?)", ids)
		if res.Error != nil || res.RowsAffected == 0 {
			return nil, err
		}
	}
	storage.DB.Model(&lesson).Updates(&models.Lesson{
		Content: lesson.Content})
	storage.DB.Where("lesson_id = ?", lesson.ID).Find(&lesson.Quizzes)
	return lesson, nil
}

func DeleteLessonByID(id string) error {
	lesson := &models.Lesson{}

	courses := []models.Course{}
	intId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("can't parse id " + id + " to int with err: " + err.Error())
	}

	query := fmt.Sprintf("SELECT * FROM courses WHERE lessons_ids @> '[%d]'", intId)
	err = storage.DB.Raw(query).Scan(&courses).Error
	if err != nil {
		return errors.New("can't query lessons from courses with err: " + err.Error())
	}

	for _, course := range courses {
		var ids []uint
		err = json.Unmarshal(course.LessonsIDs, &ids)
		if err != nil {
			return errors.New("unmarshaling ids failed with error: " + err.Error())
		}

		idx := slices.Index(ids, uint(intId))
		newLessonsIDs := slices.Delete(ids, idx, idx+1)

		if len(newLessonsIDs) == 0 {
			courseService.DeleteCourseByID(strconv.FormatUint(uint64(course.ID), 10))
		} else {
			updateCourseLessonsIDs := models.UpdateCourse{LessonsIDs: newLessonsIDs}
			err = courseService.UpdateCourseByID(strconv.FormatUint(uint64(course.ID), 10), updateCourseLessonsIDs)
			if err != nil {
				return err
			}
		}
	}
	res := storage.DB.Unscoped().Delete(lesson, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete lesson with id " + id)
	}
	return nil
}

func UpdateLessonByID(id string, updateLesson models.UpdateLesson) error {
	lesson := &models.Lesson{}
	res := storage.DB.Where("id = ?", id).Find(lesson)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the lesson, check if ID " + id + " exists")
	}

	var contentIDsJSON []byte
	if updateLesson.ContentIds != nil {
		contentIDsJSON, _ = json.Marshal(updateLesson.ContentIds)
		res := storage.DB.Debug().Find(&lesson.Content, "id IN (?)", updateLesson.ContentIds)
		if res.Error != nil || res.RowsAffected == 0 || len(lesson.Content) != len(updateLesson.ContentIds) {
			return errors.New("error in finding options when updating lesson: " +
				fmt.Sprint(len(lesson.Content)) + " lessons found, but " +
				fmt.Sprint(len(updateLesson.ContentIds)) + " lesson IDs given")
		}
	}
	storage.DB.Model(&lesson).Updates(&models.Lesson{
		Title:    updateLesson.Title,
		Order:    updateLesson.Order,
		CourseID: updateLesson.CourseID,
		Summary:  updateLesson.Summary})

	if contentIDsJSON != nil {
		storage.DB.Model(&lesson).Updates(&models.Lesson{
			ContentIds: contentIDsJSON,
			Content:    lesson.Content})
	}
	return nil
}

func AddContentsToLesson(id string, addedContentsIDs models.AddContentsToLesson) error {
	lesson := &models.Lesson{}
	res := storage.DB.Where("id = ?", id).Find(lesson)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the lesson, check if ID " + id + " exists")
	}
	if addedContentsIDs.ContentIDs == nil {
		return errors.New("no new options to append to lesson " + id)
	}
	var existingContents []uint
	err := json.Unmarshal(lesson.ContentIds, &existingContents)
	if err != nil {
		return errors.New("can't unmarshal existing option ids for lesson " + id)
	}
	var addedContents []models.Media
	res = storage.DB.Debug().Find(&addedContents, "id IN (?)", addedContentsIDs.ContentIDs)
	if res.Error != nil || res.RowsAffected == 0 || len(addedContents) != len(addedContentsIDs.ContentIDs) {
		return errors.New("error in finding medias when appending to lesson: " +
			fmt.Sprint(len(addedContents)) + " medias found, but " +
			fmt.Sprint(len(addedContentsIDs.ContentIDs)) + " medias IDs given")
	}

	for _, media := range addedContentsIDs.ContentIDs {
		idx := slices.Index(existingContents, media)
		if idx != -1 {
			return errors.New("can't have option with id " + fmt.Sprint(media) + " more than one time for lesson " + id)
		}
		existingContents = append(existingContents, media)
	}

	optionsToJSON, _ := json.Marshal(existingContents)
	storage.DB.Model(&lesson).Updates(&models.Lesson{
		ContentIds: optionsToJSON})
	return nil
}

func CreateLesson(lesson *models.Lesson) (*models.Lesson, error) {
	res := storage.DB.Model(models.Course{}).Find(&lesson.Course, "id = ?", lesson.CourseID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding course with id " + fmt.Sprint(lesson.CourseID) + " when creating lesson")
	}
	err := storage.DB.Create(lesson).Error
	if err != nil {
		return nil, err
	}
	return lesson, nil
}
