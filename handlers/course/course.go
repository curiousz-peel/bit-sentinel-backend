package handlers

import (
	"fmt"
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/course"
	"github.com/gofiber/fiber/v2"
)

func GetCourses(ctx *fiber.Ctx) error {
	courses, err := service.GetCourses()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch courses",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "courses fetched successfully",
		"data":    courses})
}

func GetCourseByID(ctx *fiber.Ctx) error {
	id := ctx.Params("courseId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "course ID cannot be empty on get",
			"data":    nil})
	}
	course, err := service.GetCourseByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the course, check if ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "course fetched successfully",
		"data":    course})
}

func DeleteCourseByID(ctx *fiber.Ctx) error {
	id := ctx.Params("courseId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "course ID cannot be empty on delete",
			"data":    nil})
	}

	err := service.DeleteCourseByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete course, check if course with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "course deleted successfully"})
}

func CreateCourse(ctx *fiber.Ctx) error {
	course := &models.Course{}
	err := ctx.BodyParser(course)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err.Error()})
		return err
	}
	courseDTO, err := service.CreateCourse(course)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create course",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "course created successfully",
		"data":    courseDTO})
}

func UpdateCourseByID(ctx *fiber.Ctx) error {
	id := ctx.Params("courseId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "course ID cannot be empty on update",
			"data":    nil})
	}

	var updateCourseData models.UpdateCourse
	err := ctx.BodyParser(&updateCourseData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: check the JSON fields",
			"data":    err.Error()})
	}

	err = service.UpdateCourseByID(id, updateCourseData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the course",
			"data":    err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "course updated successfully",
		"data":    nil})
}

func AddAuthorsToCourse(ctx *fiber.Ctx) error {
	id := ctx.Params("courseId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "course ID cannot be empty on add option",
			"data":    nil})
	}

	var addAuthorsData models.AddAuthorsToCourse
	err := ctx.BodyParser(&addAuthorsData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only add author ids to a question's options list",
			"data":    err.Error()})
	}

	err = service.AddAuthorsToCourse(id, addAuthorsData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not add authors to the question",
			"data":    err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "authors added to question successfully",
		"data":    nil})
}

func GetCoursesByRatingForHome(ctx *fiber.Ctx) error {
	courses, err := service.GetCoursesByRatingForHome()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch top 3 courses by rating",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "courses fetched successfully",
		"data":    courses})
}

func GetCoursesByMostRecentForHome(ctx *fiber.Ctx) error {
	courses, err := service.GetCoursesByMostRecentForHome()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch most recent 3 courses",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "courses fetched successfully",
		"data":    courses})
}

func GetCoursesFundamentalsForHome(ctx *fiber.Ctx) error {
	courses, err := service.GetCoursesFundamentalsForHome()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch first 3 fundamental courses",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "courses fetched successfully",
		"data":    courses})
}
