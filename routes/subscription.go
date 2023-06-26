package routes

import (
	subscriptionHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/subscription"
	"github.com/gofiber/fiber/v2"
)

func SetupSubscriptionRoutes(router fiber.Router) {
	user := router.Group("/subscription")
	user.Post("/", subscriptionHandler.CreateSubscription)
	user.Get("/", subscriptionHandler.GetSubscriptions)
	user.Get("/:subscriptionId", subscriptionHandler.GetSubscriptionByID)
	user.Put("/:subscriptionId", subscriptionHandler.UpdateSubscriptionByID)
	user.Delete("/:subscriptionId", subscriptionHandler.DeleteSubscriptionByID)
}
