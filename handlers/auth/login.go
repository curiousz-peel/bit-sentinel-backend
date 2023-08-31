package handlers

import (
	"net/http"

	jwtUtil "github.com/curiousz-peel/web-learning-platform-backend/jwt"
	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/user"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	loginUser := &models.LoginStruct{}
	err := ctx.BodyParser(loginUser)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err.Error()})
	}
	user, err := service.GetUserByUsername(loginUser.UserName)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "failed to find a user",
			"data":    err.Error()})
	}
	if user.Password != loginUser.Password {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "wrong username + password combination",
			"data":    nil})
	}
	token, err := jwtUtil.GetLoginToken(loginUser.UserName)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "could not generate token",
			"data":    err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "logged in successfully",
		"data":    token})
}
