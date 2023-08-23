package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/subscription"

	"github.com/gofiber/fiber/v2"
)

func GetSubscriptions(ctx *fiber.Ctx) error {
	subscriptions, err := service.GetSubscriptions()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch subscriptions",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscriptions fetched successfully",
		"data":    subscriptions})
}

func GetSubscriptionByID(ctx *fiber.Ctx) error {
	id := ctx.Params("subscriptionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription ID cannot be empty on get",
			"data":    nil})
	}
	subscription, err := service.GetSubscriptionByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the subscription, check if ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription fetched successfully",
		"data":    subscription})
}

func DeleteSubscriptionByID(ctx *fiber.Ctx) error {
	id := ctx.Params("subscriptionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription ID cannot be empty on delete",
			"data":    nil})
	}

	err := service.DeleteSubscriptionByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete subscription, check if subscription with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "subscription deleted successfully"})
}

func CreateSubscription(ctx *fiber.Ctx) error {
	subscription := &models.Subscription{}
	err := ctx.BodyParser(subscription)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err})
		return err
	}

	subscription, err = service.CreateSubscription(subscription)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create subscription",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription created successfully",
		"data":    subscription})
}

func UpdateSubscriptionByID(ctx *fiber.Ctx) error {
	id := ctx.Params("subscriptionId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription ID cannot be empty on update",
			"data":    nil})
	}

	var updateSubscriptionData models.UpdateSubscription
	err := ctx.BodyParser(&updateSubscriptionData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the type, duration or price",
			"data":    err})
	}

	err = service.UpdateSubscriptionByID(id, updateSubscriptionData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the subscription",
			"data":    err})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription updated successfully",
		"data":    updateSubscriptionData})
}
