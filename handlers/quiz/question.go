package handlers

import (
	"fmt"
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/quiz"
	"github.com/gofiber/fiber/v2"
)

func GetQuestions(ctx *fiber.Ctx) error {
	questions, err := service.GetQuestions()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch questions",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "questions fetched successfully",
		"data":    questions})
}

func GetQuestionByID(ctx *fiber.Ctx) error {
	id := ctx.Params("questionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "question ID cannot be empty on get",
			"data":    nil})
	}
	question, err := service.GetQuestionByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the question, check if ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "question fetched successfully",
		"data":    question})
}

func DeleteQuestionByID(ctx *fiber.Ctx) error {
	id := ctx.Params("questionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "question ID cannot be empty on delete",
			"data":    nil})
	}

	err := service.DeleteQuestionByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete question, check if question with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "question deleted successfully"})
}

func CreateQuestion(ctx *fiber.Ctx) error {
	question := &models.Question{}
	err := ctx.BodyParser(question)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err.Error()})
		return err
	}
	question, err = service.CreateQuestion(question)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create question",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "question created successfully",
		"data":    question})
}

func UpdateQuestionByID(ctx *fiber.Ctx) error {
	id := ctx.Params("questionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "question ID cannot be empty on update",
			"data":    nil})
	}

	var updateQuestionData models.UpdateQuestion
	err := ctx.BodyParser(&updateQuestionData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the question id, the text or wheter the options for that question",
			"data":    err.Error()})
	}

	err = service.UpdateQuestionByID(id, updateQuestionData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the question",
			"data":    err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "question updated successfully",
		"data":    nil})
}

func AddOptionsToQuestion(ctx *fiber.Ctx) error {
	id := ctx.Params("questionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "question ID cannot be empty on add option",
			"data":    nil})
	}

	var addOptionsData models.AddOptionsToQuestion
	err := ctx.BodyParser(&addOptionsData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only add option ids to a question's options list",
			"data":    err.Error()})
	}

	err = service.AddOptionsToQuestion(id, addOptionsData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not add options to the question",
			"data":    err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "options added to question successfully",
		"data":    nil})
}
