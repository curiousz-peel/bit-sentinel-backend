package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/handlers"
	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetSubscriptions(ctx *fiber.Ctx) error {
	return handlers.GetRecords(ctx, &[]models.Subscription{})
}

func CreateSubscription(ctx *fiber.Ctx) error {
	return handlers.CreateRecord(ctx, &models.Subscription{})
}

func GetSubscriptionByID(ctx *fiber.Ctx) error {
	return handlers.GetRecordByID(ctx, &models.Subscription{}, "subscriptionId")
}

func DeleteSubscriptionByID(ctx *fiber.Ctx) error {
	return handlers.DeleteRecordByID(ctx, &models.Subscription{}, "subscriptionId")
}

func UpdateSubscriptionByID(ctx *fiber.Ctx) error {
	type updateSubscription struct {
		Type     string  `json:"type"`
		Duration int     `json:"duration"`
		Price    float32 `json:"price"`
	}
	return handlers.UpdateRecordByID(ctx, &models.Subscription{}, &updateSubscription{}, "subscriptionId")
}

func GetSubscriptionPlans(ctx *fiber.Ctx) error {
	subscriptionPlans := &[]models.SubscriptionPlan{}
	err := storage.DB.Model(&models.SubscriptionPlan{}).Preload("User").Preload("Subscription").Find(&subscriptionPlans).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch subscription plans",
			"data":    err})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription plans fetched successfully",
		"data":    subscriptionPlans})
	return nil
}

func CreateSubscriptionPlans(ctx *fiber.Ctx) error {
	subscriptionPlan := &models.SubscriptionPlan{}
	err := ctx.BodyParser(subscriptionPlan)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "request failed",
			"data":    err})
		return err
	}

	fmt.Println(subscriptionPlan)

	res := storage.DB.Find(&subscriptionPlan.Subscription, "id = ?", subscriptionPlan.SubscriptionID)
	if res.Error != nil {
		return errors.New("error in finding subscription: " + subscriptionPlan.Subscription.Type + " with error " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return errors.New("could not find subscription: " + subscriptionPlan.Subscription.Type + " with error " + res.Error.Error())
	}

	err = storage.DB.Create(subscriptionPlan).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create subscription plan",
			"data":    err})
		return err
	}
	return nil
}

func GetSubscriptionPlanByID(ctx *fiber.Ctx) error {
	return handlers.GetRecordByID(ctx, &models.SubscriptionPlan{}, "subscriptionPlanId")
}

func DeleteSubscriptionPlanByID(ctx *fiber.Ctx) error {
	return handlers.DeleteRecordByID(ctx, &models.SubscriptionPlan{}, "subscriptionPlanId")
}

func UpdateSubscriptionPlanByID(ctx *fiber.Ctx) error {
	type updateSubscriptionPlan struct {
		UserID         uuid.UUID `json:"userID"`
		SubscriptionID int       `json:"subscriptionID"`
	}
	return handlers.UpdateRecordByID(ctx, &models.SubscriptionPlan{}, &updateSubscriptionPlan{}, "subscriptionPlanId")
}
