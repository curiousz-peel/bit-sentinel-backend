package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/subscription"
	"github.com/gofiber/fiber/v2"
)

func GetSubscriptionPlans(ctx *fiber.Ctx) error {
	subscriptionPlans, err := service.GetSubscriptionPlans()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch subscription plans",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription plans fetched successfully",
		"data":    subscriptionPlans})
}

func CreateSubscriptionPlans(ctx *fiber.Ctx) error {
	subscriptionPlan := &models.SubscriptionPlan{}
	err := ctx.BodyParser(subscriptionPlan)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err.Error()})
	}

	subscriptionPlanDTO, err := service.CreateSubscriptionPlan(subscriptionPlan)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create subscription plan",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription plan created successfully",
		"data":    subscriptionPlanDTO})
}

func GetSubscriptionPlanByID(ctx *fiber.Ctx) error {
	id := ctx.Params("subscriptionPlanId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription plan ID cannot be empty on get",
			"data":    nil})
	}
	subscriptionPlan, err := service.GetSubscriptionPlanByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the subscription plan, check if ID " + id + " exists",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription plan found successfully",
		"data":    subscriptionPlan})
}

func GetSubscriptionPlanByUserId(ctx *fiber.Ctx) error {
	id := ctx.Params("userId")
	subscriptionPlan, err := service.GetSubscriptionPlanByUserId(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch subscription plan by user id",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "subscription plan fetched successfully",
		"data":    subscriptionPlan})
}

func DeleteSubscriptionPlanByID(ctx *fiber.Ctx) error {
	id := ctx.Params("subscriptionPlanId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription plan ID cannot be empty on delete",
			"data":    nil})
	}
	err := service.DeleteSubscriptionPlanByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete subscription plan",
			"data":    err.Error()})

	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "subscription plan deleted successfully"})
}

func UpdateSubscriptionPlanByID(ctx *fiber.Ctx) error {
	id := ctx.Params("subscriptionPlanId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "subscription plan ID cannot be empty on update",
			"data":    nil})
	}
	var updateSubscriptionPlanData models.UpdateSubscriptionPlan
	err := ctx.BodyParser(&updateSubscriptionPlanData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the user id or the subscription type id",
			"data":    err.Error()})
	}
	err = service.UpdateSubscriptionPlanByID(id, updateSubscriptionPlanData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the subscription plan",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "comment updated successfully",
		"data":    nil})
}
