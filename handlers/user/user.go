package handlers

import (
	"github.com/curiousz-peel/web-learning-platform-backend/handlers"
	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {
	return handlers.GetRecords(ctx, &[]models.User{})
}

func CreateUser(ctx *fiber.Ctx) error {
	return handlers.CreateRecord(ctx, &models.User{})
}

func GetUserByID(ctx *fiber.Ctx) error {
	return handlers.GetRecordByID(ctx, &models.User{}, "userId")
}

func DeleteUserByID(ctx *fiber.Ctx) error {
	return handlers.DeleteRecordByID(ctx, &models.User{}, "userId")
}

func UpdateUserByID(ctx *fiber.Ctx) error {
	type updateUser struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		IsAuthor  string `json:"isAuthor"`
	}
	return handlers.UpdateRecordByID(ctx, &models.User{}, &updateUser{}, "userId")
}
