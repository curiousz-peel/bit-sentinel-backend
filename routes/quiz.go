package routes

import (
	quizHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/quiz"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func SetupQuizRoutes(router fiber.Router) {
	quiz := router.Group("/quiz")
	quiz.Post("/", jwtHandler.ValidateToken, quizHandler.CreateQuiz)
	quiz.Get("/", jwtHandler.ValidateToken, quizHandler.GetQuizzes)
	quiz.Get("/:quizId", jwtHandler.ValidateToken, quizHandler.GetQuizByID)
	quiz.Put("/:quizId", jwtHandler.ValidateToken, quizHandler.UpdateQuizByID)
	quiz.Put("addQuestions/:quizId", jwtHandler.ValidateToken, quizHandler.AddQuestionsToQuiz)
	quiz.Delete("/:quizId", jwtHandler.ValidateToken, quizHandler.DeleteQuizByID)
}

func SetupQuestionRoutes(router fiber.Router) {
	question := router.Group("/question")
	question.Post("/", jwtHandler.ValidateToken, quizHandler.CreateQuestion)
	question.Get("/", jwtHandler.ValidateToken, quizHandler.GetQuestions)
	question.Get("/:questionId", jwtHandler.ValidateToken, quizHandler.GetQuestionByID)
	question.Put("/:questionId", jwtHandler.ValidateToken, quizHandler.UpdateQuestionByID)
	question.Put("addOptions/:questionId", jwtHandler.ValidateToken, quizHandler.AddOptionsToQuestion)
	question.Delete("/:questionId", jwtHandler.ValidateToken, quizHandler.DeleteQuestionByID)
}

func SetupOptionRoutes(router fiber.Router) {
	option := router.Group("/option")
	option.Post("/", jwtHandler.ValidateToken, quizHandler.CreateOption)
	option.Get("/", jwtHandler.ValidateToken, quizHandler.GetOptions)
	option.Get("/:optionId", jwtHandler.ValidateToken, quizHandler.GetOptionByID)
	option.Put("/:optionId", jwtHandler.ValidateToken, quizHandler.UpdateOptionByID)
	option.Delete("/:optionId", jwtHandler.ValidateToken, quizHandler.DeleteOptionByID)
}

func SetupQuizRelatedRoutes(router fiber.Router) {
	SetupQuizRoutes(router)
	SetupQuestionRoutes(router)
	SetupOptionRoutes(router)
}
