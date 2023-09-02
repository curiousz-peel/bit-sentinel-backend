package handlers

import (
	"net/http"
	"time"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	jwtUtil "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/user"

	// jwtUtil "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func Signup(ctx *fiber.Ctx) error {
	user := &models.User{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err.Error()})
	}
	bday, err := time.Parse("2006-01-02T15:04:05", "1999-04-23T12:34:56")
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "bday stuff :c",
			"data":    err.Error()})
	}
	user.Birthday = bday

	_, err = service.CreateUser(user)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "failed to create an account",
			"data":    err.Error()})
	}

	token, err := jwtUtil.GetLoginToken(user.UserName)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "could not generate token",
			"data":    err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "account creation succeeded",
		"data":    token})
}
