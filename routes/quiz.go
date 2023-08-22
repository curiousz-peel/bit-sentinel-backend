package routes

import (
	quizHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/quiz"
	"github.com/gofiber/fiber/v2"
)

func SetupQuizRoutes(router fiber.Router) {
	quiz := router.Group("/quiz")
	quiz.Post("/", quizHandler.CreateQuiz)
	quiz.Get("/", quizHandler.GetQuizzes)
	quiz.Get("/:quizId", quizHandler.GetQuizByID)
	quiz.Put("/:quizId", quizHandler.UpdateQuizByID)
	quiz.Delete("/:quizId", quizHandler.DeleteQuizByID)
}

func SetupQuestionRoutes(router fiber.Router) {
	question := router.Group("/question")
	question.Post("/", quizHandler.CreateQuestion)
	question.Get("/", quizHandler.GetQuestions)
	question.Get("/:questionId", quizHandler.GetQuestionByID)
	question.Put("/:questionId", quizHandler.UpdateQuestionByID)
	question.Put("addOptions/:questionId", quizHandler.AddOptionsToQuestion)
	question.Delete("/:questionId", quizHandler.DeleteQuestionByID)
}

func SetupOptionRoutes(router fiber.Router) {
	option := router.Group("/option")
	option.Post("/", quizHandler.CreateOption)
	option.Get("/", quizHandler.GetOptions)
	option.Get("/:optionId", quizHandler.GetOptionByID)
	option.Put("/:optionId", quizHandler.UpdateOptionByID)
	option.Delete("/:optionId", quizHandler.DeleteOptionByID)
}

func SetupQuizRelatedRoutes(router fiber.Router) {
	SetupQuizRoutes(router)
	SetupQuestionRoutes(router)
	SetupOptionRoutes(router)
}
