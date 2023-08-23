package handlers

import (
	"fmt"
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/progress"

	"github.com/gofiber/fiber/v2"
)

func GetProgresss(ctx *fiber.Ctx) error {
	progresss, err := service.GetProgresss()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch progresss",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "progresss fetched successfully",
		"data":    progresss})
}

func GetProgressByID(ctx *fiber.Ctx) error {
	id := ctx.Params("progressId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "progress ID cannot be empty on get",
			"data":    nil})
	}
	progress, err := service.GetProgressByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the progress, check if ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "progress fetched successfully",
		"data":    progress})
}

func DeleteProgressByID(ctx *fiber.Ctx) error {
	id := ctx.Params("progressId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "progress ID cannot be empty on delete",
			"data":    nil})
	}

	err := service.DeleteProgressByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete progress, check if progress with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "progress deleted successfully"})
}

func CreateProgress(ctx *fiber.Ctx) error {
	progress := &models.Progress{}
	err := ctx.BodyParser(progress)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err})
		return err
	}

	progress, err = service.CreateProgress(progress)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create progress",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "progress created successfully",
		"data":    progress})
}

func UpdateProgressByID(ctx *fiber.Ctx) error {
	id := ctx.Params("progressId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "progress ID cannot be empty on update",
			"data":    nil})
	}

	var updateProgressData models.UpdateProgress
	err := ctx.BodyParser(&updateProgressData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the user id, the course id or the text",
			"data":    err})
	}

	err = service.UpdateProgressByID(id, updateProgressData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the progress",
			"data":    err})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "progress updated successfully",
		"data":    updateProgressData})
}
