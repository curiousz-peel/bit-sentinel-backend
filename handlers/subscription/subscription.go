package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

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

	res := storage.DB.Find(&subscriptionPlan.Subscription, "id = ?", subscriptionPlan.SubscriptionID)
	if res.Error != nil {
		return errors.New("error in finding subscription: " + subscriptionPlan.Subscription.Type + " with error " + res.Error.Error())
	} else if res.RowsAffected == 0 {
		return errors.New("could not find subscription: " + subscriptionPlan.Subscription.Type + " with error " + res.Error.Error())
	}

	subscriptionPlan.CreatedAt = time.Now()
	subscriptionPlan.StartDate = time.Now()

	if subscriptionPlan.SubscriptionID != 1 {
		subscriptionPlan.EndDate = time.Now().Add(time.Hour * 24 * time.Duration(subscriptionPlan.Subscription.Duration))
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
	subscriptionPlan := &models.SubscriptionPlan{}
	id := ctx.Params("subscriptionPlanId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription plan ID cannot be empty on get",
			"data":    nil})
	}

	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Subscription").Find(subscriptionPlan)
	if res.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the subscription plan",
			"data":    res.Error})
		return res.Error
	} else if res.RowsAffected == 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "could not find subscription plan with id: " + id + " to fetch"})
		return errors.New("could not find subscription plan with id: " + id + " to fetch")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription plan found successfully",
		"data":    subscriptionPlan})
	return nil
}

func DeleteSubscriptionPlanByID(ctx *fiber.Ctx) error {
	subscriptionPlan := &models.SubscriptionPlan{}
	id := ctx.Params("subscriptionPlanId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription plan ID cannot be empty on delete",
			"data":    nil})
	}

	res := storage.DB.Unscoped().Delete(subscriptionPlan, id)
	if res.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete subscription plan",
			"data":    res.Error})
		return res.Error
	} else if res.RowsAffected == 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "could not find a subscription plan with id: " + id + " to delete"})
		return nil
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "subscription plan deleted successfully"})
	return nil
}

func UpdateSubscriptionPlanByID(ctx *fiber.Ctx) error {
	type updateSubscriptionPlan struct {
		UserID         uuid.UUID `json:"userID"`
		SubscriptionID int       `json:"subscriptionID"`
	}
	subscriptionPlan := &models.SubscriptionPlan{}
	id := ctx.Params("subscriptionPlanId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription plan ID cannot be empty on update",
			"data":    nil})
	}

	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Subscription").Find(subscriptionPlan)
	if res.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the subscription plan",
			"data":    res.Error})
		return res.Error
	} else if res.RowsAffected == 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "could not find the subscription plan"})
		return nil
	}

	var updateSubscriptionPlanData updateSubscriptionPlan
	err := ctx.BodyParser(&updateSubscriptionPlanData)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the user id or the subscription type id",
			"data":    err})
		return err
	}

	storage.DB.Model(&subscriptionPlan).Updates(&models.SubscriptionPlan{
		SubscriptionID: uint(updateSubscriptionPlanData.SubscriptionID),
		UserID:         updateSubscriptionPlanData.UserID})

	type UpdateSubscriptionPlanResponse struct {
		SubscriptionID uint
		UserID         string
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription plan updated successfully",
		"data":    UpdateSubscriptionPlanResponse{SubscriptionID: subscriptionPlan.SubscriptionID, UserID: subscriptionPlan.UserID.String()}})
	return nil
}
