package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/user"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {
	users, err := service.GetUsers()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch users",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "users fetched successfully",
		"data":    users})
}

func CreateUser(ctx *fiber.Ctx) error {
	user := &models.User{}
	err := ctx.BodyParser(user)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err})
	}
	user, err = service.CreateUser(user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create user",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create user succeeded",
		"data":    user})
}

func GetUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("userId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "user ID cannot be empty on get",
			"data":    nil})
	}
	user, err := service.GetUserByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the user, check if ID " + id + " exists",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user fetched successfully",
		"data":    user})
}

func DeleteUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("userId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "user ID cannot be empty on delete",
			"data":    nil})
	}
	err := service.DeleteUserByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete user, check if user with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "user deleted successfully"})
}

func UpdateUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("userId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "user ID cannot be empty on update",
			"data":    nil})
	}
	var updateUserData models.UpdateUser
	err := ctx.BodyParser(&updateUserData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the first name, last name, user name, email, password, birthday, or is mod status",
			"data":    err.Error()})
	}
	err = service.UpdateUserByID(id, updateUserData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the user",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user updated successfully",
		"data":    nil})
}
