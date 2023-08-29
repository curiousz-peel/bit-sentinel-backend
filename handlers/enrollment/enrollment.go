package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/enrollment"
	"github.com/gofiber/fiber/v2"
)

func GetEnrollments(ctx *fiber.Ctx) error {
	enrollments, err := service.GetEnrollments()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch enrollments",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "enrollments fetched successfully",
		"data":    enrollments})
}

func GetEnrollmentByID(ctx *fiber.Ctx) error {
	id := ctx.Params("enrollmentId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "enrollment ID cannot be empty on get",
			"data":    nil})
	}
	enrollment, err := service.GetEnrollmentByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the enrollment, check if ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "enrollment fetched successfully",
		"data":    enrollment})
}

func DeleteEnrollmentByID(ctx *fiber.Ctx) error {
	id := ctx.Params("enrollmentId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "enrollment ID cannot be empty on delete",
			"data":    nil})
	}

	err := service.DeleteEnrollmentByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete enrollment, check if enrollment with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "enrollment deleted successfully"})
}

func CreateEnrollment(ctx *fiber.Ctx) error {
	enrollment := &models.Enrollment{}
	err := ctx.BodyParser(enrollment)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err})
	}

	enrollmentDTO, err := service.CreateEnrollment(enrollment)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to create enrollment",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "enrollment created successfully",
		"data": enrollmentDTO})
}

func UpdateEnrollmentByID(ctx *fiber.Ctx) error {
	id := ctx.Params("enrollmentId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "enrollment ID cannot be empty on update",
			"data":    nil})
	}
	var updateEnrollmentData models.UpdateEnrollment
	err := ctx.BodyParser(&updateEnrollmentData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the user id, the course id or the progress id",
			"data":    err})
	}
	err = service.UpdateEnrollmentByID(id, updateEnrollmentData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to updated enrollment",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "enrollment updated successfully",
		"data":    updateEnrollmentData})
}
