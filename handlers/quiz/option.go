package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/quiz"
	"github.com/gofiber/fiber/v2"
)

func GetOptions(ctx *fiber.Ctx) error {
	options, err := service.GetOptions()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch options",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "options fetched successfully",
		"data":    options})
}

func GetOptionByID(ctx *fiber.Ctx) error {
	id := ctx.Params("optionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "option ID cannot be empty on get",
			"data":    nil})
	}
	option, err := service.GetOptionByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the option, check if ID " + id + " exists",
			"data":    nil})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "option fetched successfully",
		"data":    option})
}

func DeleteOptionByID(ctx *fiber.Ctx) error {
	id := ctx.Params("optionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "option ID cannot be empty on delete",
			"data":    nil})
	}

	err := service.DeleteOptionByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete option, check if option with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "option deleted successfully"})
}

func CreateOption(ctx *fiber.Ctx) error {
	option := &models.Option{}
	err := ctx.BodyParser(option)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err})
		return err
	}

	optionCreated, err := service.CreateOption(option)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create option",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "option created successfully",
		"data":    optionCreated})
}

func UpdateOptionByID(ctx *fiber.Ctx) error {
	id := ctx.Params("optionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "option ID cannot be empty on update",
			"data":    nil})
	}

	var updateOptionData models.UpdateOption
	err := ctx.BodyParser(&updateOptionData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the question id, the text or wheter the option is correct or not",
			"data":    err})
	}

	err = service.UpdateOptionByID(id, updateOptionData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the option",
			"data":    err})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "option updated successfully",
		"data":    updateOptionData})
}
