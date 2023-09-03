package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/lesson"
	"github.com/gofiber/fiber/v2"
)

func GetLessons(ctx *fiber.Ctx) error {
	lessons, err := service.GetLessons()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch lessons",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "lessons fetched successfully",
		"data":    lessons})
}

func GetLessonsByCourseId(ctx *fiber.Ctx) error {
	courseId := ctx.Params("courseId")
	if courseId == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "course ID cannot be empty on get lessons",
			"data":    nil})
	}
	lessons, err := service.GetLessonsByCourseId(courseId)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the lessons, check if course ID " + courseId + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "lesson sfetched successfully",
		"data":    lessons})
}

func GetLessonByID(ctx *fiber.Ctx) error {
	id := ctx.Params("lessonId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "lesson ID cannot be empty on get",
			"data":    nil})
	}
	lesson, err := service.GetLessonByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the lesson, check if ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "lesson fetched successfully",
		"data":    lesson})
}

func DeleteLessonByID(ctx *fiber.Ctx) error {
	id := ctx.Params("lessonId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "lesson ID cannot be empty on delete",
			"data":    nil})
	}

	err := service.DeleteLessonByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete lesson, check if lesson with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "lesson deleted successfully"})
}

func CreateLesson(ctx *fiber.Ctx) error {
	lesson := &models.Lesson{}
	err := ctx.BodyParser(lesson)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err.Error()})
		return err
	}
	lessonDTO, err := service.CreateLesson(lesson)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create lesson",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "lesson created successfully",
		"data":    lessonDTO})
}

func UpdateLessonByID(ctx *fiber.Ctx) error {
	id := ctx.Params("lessonId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "lesson ID cannot be empty on update",
			"data":    nil})
	}

	var updateLessonData models.UpdateLesson
	err := ctx.BodyParser(&updateLessonData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the lesson id, the text or wheter the options for that lesson",
			"data":    err.Error()})
	}

	err = service.UpdateLessonByID(id, updateLessonData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the lesson",
			"data":    err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "lesson updated successfully",
		"data":    nil})
}
