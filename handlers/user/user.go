package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/curiousz-peel/web-learning-platform-backend/handlers"
	subscriptionHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/subscription"
	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {
	return handlers.GetRecords(ctx, &[]models.User{})
}

// take User entity out of utils, clear mess. make automatic creation of basic subscription plan for user on user creation
func CreateUser(ctx *fiber.Ctx) error {
	user := &models.User{}
	err := ctx.BodyParser(user)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "request failed",
			"data":    err})
		return err
	}

	err = storage.DB.Create(user).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create user",
			"data":    err})
		return err
	}

	//create Basic subscription when a new user registers
	basicSubscription := map[string]interface{}{"StartDate": time.Now(),
		"EndDate":        time.Now().AddDate(1000, 0, 0),
		"SubscriptionID": 1,
		"UserID":         user.ID}

	basicSubscriptionJSON, err := json.Marshal(basicSubscription)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "user was created, but Basic subscription plan failed to be parsed",
			"data":    err})
		return err
	}

	ctx.Context().Request.ResetBody()
	ctx.Context().Request.SetBodyString(string(basicSubscriptionJSON))
	err = subscriptionHandler.CreateSubscriptionPlans(ctx)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "user was created, but creation of Basic subscription failed",
			"data":    err})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create user succeeded",
		"data":    user})
	return nil
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
