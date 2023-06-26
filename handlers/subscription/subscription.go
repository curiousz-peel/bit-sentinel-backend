package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetSubscriptions(ctx *fiber.Ctx) error {
	subscription := &[]models.Subscription{}

	err := storage.DB.Find(subscription).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch subscription types",
			"data":    err})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscriptions fetched successfully",
		"data":    subscription})
	return nil
}

func CreateSubscription(ctx *fiber.Ctx) error {
	subscription := models.Subscription{}
	err := ctx.BodyParser(&subscription)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "request failed",
			"data":    err})
		return err
	}

	err = storage.DB.Create(&subscription).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create subscription type",
			"data":    err})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create subscription succeeded",
		"data":    subscription})
	return nil
}

func GetSubscriptionByID(ctx *fiber.Ctx) error {
	subscription := &models.Subscription{}
	id := ctx.Params("subscriptionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription type ID cannot be empty on get",
			"data":    nil})
	}
	err := storage.DB.Where("id = ?", id).Find(subscription).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the subscription type",
			"data":    err})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription type found",
		"data":    subscription})
	return nil
}

func DeleteSubscriptionByID(ctx *fiber.Ctx) error {
	subscription := &models.Subscription{}
	id := ctx.Params("subscriptionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription type ID cannot be empty on delete"})
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid subscription type ID",
			"data":    err.Error(),
		})
	}

	err = storage.DB.Delete(subscription, uuid).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete subscription type",
			"data":    err})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "subscription deleted successfully"})
	return nil
}

func UpdateSubscriptionByID(ctx *fiber.Ctx) error {
	type updateSubscription struct {
		Type     string  `json:"type"`
		Duration int     `json:"duration"`
		Price    float32 `json:"price"`
	}
	subscription := &models.Subscription{}
	id := ctx.Params("subscriptionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription type ID cannot be empty on get",
			"data":    nil})
	}
	err := storage.DB.Where("id = ?", id).Find(subscription).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the subscription type",
			"data":    err})
		return err
	}

	var updateSubscriptionData updateSubscription
	err = ctx.BodyParser(&updateSubscriptionData)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the type, duration and price",
			"data":    err})
		return err
	}

	storage.DB.Model(&subscription).Updates(&models.Subscription{
		Type:     updateSubscriptionData.Type,
		Duration: updateSubscriptionData.Duration,
		Price:    updateSubscriptionData.Price})

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription updated successfully",
		"data":    subscription})
	return nil
}
