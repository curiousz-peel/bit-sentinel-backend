package routes

import (
	subscriptionHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/subscription"
	"github.com/gofiber/fiber/v2"
)

func SetupSubscriptionRoutes(router fiber.Router) {
	subscription := router.Group("/subscription")
	subscription.Post("/", subscriptionHandler.CreateSubscription)
	subscription.Get("/", subscriptionHandler.GetSubscriptions)
	subscription.Get("/:subscriptionId", subscriptionHandler.GetSubscriptionByID)
	subscription.Put("/:subscriptionId", subscriptionHandler.UpdateSubscriptionByID)
	subscription.Delete("/:subscriptionId", subscriptionHandler.DeleteSubscriptionByID)
}
