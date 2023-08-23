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

func SetupSubscriptionPlanRoutes(router fiber.Router) {
	subscriptionPlan := router.Group("/plan")
	subscriptionPlan.Post("/", subscriptionHandler.CreateSubscriptionPlans)
	subscriptionPlan.Get("/", subscriptionHandler.GetSubscriptionPlans)
	subscriptionPlan.Get("/:subscriptionPlanId", subscriptionHandler.GetSubscriptionPlanByID)
	subscriptionPlan.Put("/:subscriptionPlanId", subscriptionHandler.UpdateSubscriptionPlanByID)
	subscriptionPlan.Delete("/:subscriptionPlanId", subscriptionHandler.DeleteSubscriptionPlanByID)
}

func SetupSubscriptionRelatedRoutes(router fiber.Router) {
	SetupSubscriptionRoutes(router)
	SetupSubscriptionPlanRoutes(router)
}
