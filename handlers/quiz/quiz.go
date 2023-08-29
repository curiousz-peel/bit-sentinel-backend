package handlers

import (
	"fmt"
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/quiz"
	"github.com/gofiber/fiber/v2"
)

func GetQuizzes(ctx *fiber.Ctx) error {
	quizzes, err := service.GetQuizzes()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch quizzes",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "quizzes fetched successfully",
		"data":    quizzes})
}

func GetQuizByID(ctx *fiber.Ctx) error {
	id := ctx.Params("quizId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "quiz ID cannot be empty on get",
			"data":    nil})
	}
	quiz, err := service.GetQuizByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the quiz, check if ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "quiz fetched successfully",
		"data":    quiz})
}

func DeleteQuizByID(ctx *fiber.Ctx) error {
	id := ctx.Params("quizId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "quiz ID cannot be empty on delete",
			"data":    nil})
	}

	err := service.DeleteQuizByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete quiz, check if quiz with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "quiz deleted successfully"})
}

func CreateQuiz(ctx *fiber.Ctx) error {
	quiz := &models.Quiz{}
	err := ctx.BodyParser(quiz)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err})
		return err
	}
	quizDTO, err := service.CreateQuiz(quiz)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create quiz",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "quiz created successfully",
		"data":    quizDTO})
}

func UpdateQuizByID(ctx *fiber.Ctx) error {
	id := ctx.Params("quizId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "quiz ID cannot be empty on update",
			"data":    nil})
	}

	var updateQuizData models.UpdateQuiz
	err := ctx.BodyParser(&updateQuizData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the quiz id, the text or wheter the quiz is correct or not",
			"data":    err.Error()})
	}

	err = service.UpdateQuizByID(id, updateQuizData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the quiz",
			"data":    err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "quiz updated successfully",
		"data":    nil})
}

func AddQuestionsToQuiz(ctx *fiber.Ctx) error {
	id := ctx.Params("quizId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "quiz ID cannot be empty on appending questions",
			"data":    nil})
	}

	var addQuestionsData models.AddQuestionsToQuiz
	err := ctx.BodyParser(&addQuestionsData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only add question ids to a quiz",
			"data":    err.Error()})
	}

	err = service.AddQuestionsToQuiz(id, addQuestionsData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not add questions to quiz",
			"data":    err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "questions added to quiz successfully",
		"data":    nil})
}
