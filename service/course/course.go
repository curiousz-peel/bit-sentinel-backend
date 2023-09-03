package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/google/uuid"
)

func GetCourses() (*[]models.CourseDTO, error) {
	courses := []models.Course{}
	coursesDTO := []models.CourseDTO{}
	err := storage.DB.Find(&courses).Error
	if err != nil {
		return nil, err
	}
	courses, err = populateCourses(courses)
	if err != nil {
		return nil, err
	}
	for _, course := range courses {
		coursesDTO = append(coursesDTO, models.ToCourseDTO(course))
	}
	return &coursesDTO, nil
}

func GetCourseByID(id string) (*models.CourseDTO, error) {
	course := &models.Course{}
	res := storage.DB.Where("id = ?", id).Find(course)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("can't find course with id " + id)
	}
	if course.AuthorsIDs != nil {
		var ids []uuid.UUID
		err := json.Unmarshal(course.AuthorsIDs, &ids)
		if err != nil {
			return nil, err
		}
		res = storage.DB.Find(&course.Authors, "id IN (?)", ids)
		if res.Error != nil || res.RowsAffected == 0 {
			return nil, errors.New("couldn't find authors for course " + fmt.Sprint(course.ID))
		}
	}
	storage.DB.Model(&course).Updates(&models.Course{
		Authors: course.Authors})
	storage.DB.Where("course_id = ?", course.ID).Find(&course.Lessons)
	storage.DB.Where("course_id = ?", course.ID).Find(&course.Ratings)
	storage.DB.Where("course_id = ?", course.ID).Find(&course.Comments)
	courseDTO := models.ToCourseDTO(*course)
	return &courseDTO, nil
}

func DeleteCourseByID(id string) error {
	course := &models.Course{}
	res := storage.DB.Unscoped().Delete(course, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete course with id " + id)
	}
	return nil
}

func UpdateCourseByID(id string, updateCourse models.UpdateCourse) error {
	course := &models.Course{}
	res := storage.DB.Where("id = ?", id).Find(course)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the course, check if ID " + id + " exists")
	}

	var authorsIDsJSON []byte
	if updateCourse.AuthorsIDs != nil {
		authorsIDsJSON, _ = json.Marshal(updateCourse.AuthorsIDs)
		res := storage.DB.Debug().Find(&course.Authors, "id IN (?)", updateCourse.AuthorsIDs)
		if res.Error != nil || res.RowsAffected == 0 || len(course.Authors) != len(updateCourse.AuthorsIDs) {
			return errors.New("error in finding authors when updating course: " +
				fmt.Sprint(len(course.Authors)) + " authors found, but " +
				fmt.Sprint(len(updateCourse.AuthorsIDs)) + " authors IDs given")
		}
	}
	storage.DB.Model(&course).Updates(&models.Course{
		Title:                 updateCourse.Title,
		Tags:                  updateCourse.Tags,
		Visible:               updateCourse.Visible,
		Rating:                updateCourse.Rating,
		IncludedSubscriptions: updateCourse.IncludedSubscriptions})

	if authorsIDsJSON != nil {
		storage.DB.Model(&course).Updates(&models.Course{
			AuthorsIDs: authorsIDsJSON,
			Authors:    course.Authors})
	}
	return nil
}

func AddAuthorsToCourse(id string, addedAuthorsIDs models.AddAuthorsToCourse) error {
	course := &models.Course{}
	res := storage.DB.Where("id = ?", id).Find(course)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the course, check if ID " + id + " exists")
	}
	if addedAuthorsIDs.AuthorsIDs == nil {
		return errors.New("no new options to append to course " + id)
	}
	var existingAuthors []uuid.UUID
	err := json.Unmarshal(course.AuthorsIDs, &existingAuthors)
	if err != nil {
		return errors.New("can't unmarshal existing authors ids for course " + id)
	}
	var addedAuthors []models.Author
	res = storage.DB.Debug().Find(&addedAuthors, "id IN (?)", addedAuthorsIDs.AuthorsIDs)
	if res.Error != nil || res.RowsAffected == 0 || len(addedAuthors) != len(addedAuthorsIDs.AuthorsIDs) {
		return errors.New("error in finding authors when appending to course: " +
			fmt.Sprint(len(addedAuthors)) + " authors found, but " +
			fmt.Sprint(len(addedAuthorsIDs.AuthorsIDs)) + " authors IDs given")
	}
	for _, author := range addedAuthorsIDs.AuthorsIDs {
		idx := slices.Index(existingAuthors, author)
		if idx != -1 {
			return errors.New("can't have option with id " + fmt.Sprint(author) + " more than one time for course " + id)
		}
		existingAuthors = append(existingAuthors, author)
	}

	authorsToJSON, _ := json.Marshal(existingAuthors)
	storage.DB.Model(&course).Updates(&models.Course{
		AuthorsIDs: authorsToJSON})
	return nil
}

func CreateCourse(course *models.Course) (*models.CourseDTO, error) {
	var authorsIds []uuid.UUID
	err := json.Unmarshal(course.AuthorsIDs, &authorsIds)
	if err != nil {
		return nil, errors.New("can't unmarshal authors ids for course " + fmt.Sprint(course.ID))
	}

	res := storage.DB.Find(&course.Authors, "id IN (?)", authorsIds)
	if res.Error != nil || res.RowsAffected == 0 || len(course.Authors) != len(authorsIds) {
		return nil, errors.New("error in finding authors when creating course: " +
			fmt.Sprint(len(course.Authors)) + " authors found, but " +
			fmt.Sprint(len(authorsIds)) + " authors IDs given")
	}
	err = storage.DB.Create(course).Error
	if err != nil {
		return nil, err
	}
	courseDTO := models.ToCourseDTO(*course)
	return &courseDTO, nil
}

func GetCoursesByRatingForHome() (*[]models.CourseDTO, error) {
	courses := []models.Course{}
	coursesDTO := []models.CourseDTO{}
	err := storage.DB.Order("rating desc").Limit(3).Find(&courses).Error
	if err != nil {
		return nil, err
	}
	courses, err = populateCourses(courses)
	if err != nil {
		return nil, err
	}
	for _, course := range courses {
		coursesDTO = append(coursesDTO, models.ToCourseDTO(course))
	}
	return &coursesDTO, nil
}

func GetCoursesByMostRecentForHome() (*[]models.CourseDTO, error) {
	courses := []models.Course{}
	coursesDTO := []models.CourseDTO{}
	err := storage.DB.Order("created_at desc").Limit(3).Find(&courses).Error
	if err != nil {
		return nil, err
	}
	courses, err = populateCourses(courses)
	if err != nil {
		return nil, err
	}
	for _, course := range courses {
		coursesDTO = append(coursesDTO, models.ToCourseDTO(course))
	}
	return &coursesDTO, nil
}

func GetCoursesFundamentalsForHome() (*[]models.CourseDTO, error) {
	courses := []models.Course{}
	coursesDTO := []models.CourseDTO{}
	// err := storage.DB.Limit(3).Find(&courses, "included_subscriptions ? 'Basic'").Error
	err := storage.DB.Limit(3).Find(&courses, "tags ? 'tag1'").Error
	if err != nil {
		return nil, err
	}
	courses, err = populateCourses(courses)
	if err != nil {
		return nil, err
	}
	for _, course := range courses {
		coursesDTO = append(coursesDTO, models.ToCourseDTO(course))
	}
	return &coursesDTO, nil
}

func GetCoursesBySubscription(subscriptionType string) (*[]models.CourseDTO, error) {
	courses := []models.Course{}
	coursesDTO := []models.CourseDTO{}
	sqlCondition := "included_subscriptions ? '" + subscriptionType + "'"
	err := storage.DB.Find(&courses, sqlCondition).Error
	if err != nil {
		return nil, err
	}
	courses, err = populateCourses(courses)
	if err != nil {
		return nil, err
	}
	for _, course := range courses {
		coursesDTO = append(coursesDTO, models.ToCourseDTO(course))
	}
	return &coursesDTO, nil
}

func GetCoursesByAuthorId(id string) (*[]models.CourseDTO, error) {
	courses := []models.Course{}
	coursesDTO := []models.CourseDTO{}
	sqlCondition := "authors_ids ? '" + id + "'"
	err := storage.DB.Find(&courses, sqlCondition).Error
	if err != nil {
		return nil, errors.New("error in finding courses by " + id)
	}
	courses, err = populateCourses(courses)
	if err != nil {
		return nil, err
	}
	for _, course := range courses {
		coursesDTO = append(coursesDTO, models.ToCourseDTO(course))
	}
	return &coursesDTO, nil
}

func populateCourses(courses []models.Course) ([]models.Course, error) {
	for i := range courses {
		if courses[i].AuthorsIDs != nil {
			var ids []uuid.UUID
			err := json.Unmarshal(courses[i].AuthorsIDs, &ids)
			if err != nil {
				return nil, err
			}
			if len(ids) > 0 {
				res := storage.DB.Find(&courses[i].Authors, "id IN (?)", ids)
				if res.Error != nil || res.RowsAffected == 0 {
					return nil, errors.New("can't load authors for course with id " + fmt.Sprint(courses[i].ID))
				}
			}
		}
		storage.DB.Where("course_id = ?", courses[i].ID).Find(&courses[i].Lessons)
		storage.DB.Where("course_id = ?", courses[i].ID).Find(&courses[i].Ratings)
		storage.DB.Where("course_id = ?", courses[i].ID).Find(&courses[i].Comments)
	}
	return courses, nil
}
