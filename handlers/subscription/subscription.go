package handlers

import (
	"github.com/curiousz-peel/web-learning-platform-backend/handlers"
	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/gofiber/fiber/v2"
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
