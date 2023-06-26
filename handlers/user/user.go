package handlers

import (
	"fmt"
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetUsers(ctx *fiber.Ctx) error {
	users := &[]models.User{}

	err := storage.DB.Find(users).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch users",
			"data":    err})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "users fetched successfully",
		"data":    users})
	return nil
}

func CreateUser(ctx *fiber.Ctx) error {
	user := models.User{}
	err := ctx.BodyParser(&user)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "request failed",
			"data":    err})
		return err
	}

	err = storage.DB.Create(&user).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create user",
			"data":    err})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create user succeeded",
		"data":    user})
	return nil
}

func GetUserByID(ctx *fiber.Ctx) error {
	user := &models.User{}
	id := ctx.Params("userId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "user ID cannot be empty on get",
			"data":    nil})
	}
	fmt.Println("the user id is", id)
	err := storage.DB.Where("id = ?", id).Find(user).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the user",
			"data":    err})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user found successfully",
		"data":    user})
	return nil
}

func DeleteUserByID(ctx *fiber.Ctx) error {
	user := &models.User{}
	id := ctx.Params("userId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "user ID cannot be empty on delete"})
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid user ID",
			"data":    err.Error(),
		})
	}

	err = storage.DB.Delete(user, uuid).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete user",
			"data":    err})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "user deleted successfully"})
	return nil
}

func UpdateUserByID(ctx *fiber.Ctx) error {
	type updateUser struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
	user := &models.User{}
	id := ctx.Params("userId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "user ID cannot be empty on get",
			"data":    nil})
	}
	fmt.Println("the user id is", id)
	err := storage.DB.Where("id = ?", id).Find(user).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the user",
			"data":    err})
		return err
	}

	var updateUserData updateUser
	err = ctx.BodyParser(&updateUserData)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update first name, last name, email or password fields",
			"data":    err})
		return err
	}

	storage.DB.Model(&user).Updates(&models.User{
		FirstName: updateUserData.FirstName,
		LastName:  updateUserData.LastName,
		Email:     updateUserData.Email,
		Password:  updateUserData.Password})

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user updated successfully",
		"data":    user})
	return nil
}
