package routes

import (
	subscriptionHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/subscription"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func SetupSubscriptionRoutes(router fiber.Router) {
	subscription := router.Group("/subscription")
	subscription.Post("/", jwtHandler.ValidateToken, subscriptionHandler.CreateSubscription)
	subscription.Get("/", jwtHandler.ValidateToken, subscriptionHandler.GetSubscriptions)
	subscription.Get("/:subscriptionId", jwtHandler.ValidateToken, subscriptionHandler.GetSubscriptionByID)
	subscription.Put("/:subscriptionId", jwtHandler.ValidateToken, subscriptionHandler.UpdateSubscriptionByID)
	subscription.Delete("/:subscriptionId", jwtHandler.ValidateToken, subscriptionHandler.DeleteSubscriptionByID)
}

func SetupSubscriptionPlanRoutes(router fiber.Router) {
	subscriptionPlan := router.Group("/plan")
	subscriptionPlan.Post("/", jwtHandler.ValidateToken, subscriptionHandler.CreateSubscriptionPlans)
	subscriptionPlan.Get("/", jwtHandler.ValidateToken, subscriptionHandler.GetSubscriptionPlans)
	subscriptionPlan.Get("/:subscriptionPlanId", jwtHandler.ValidateToken, subscriptionHandler.GetSubscriptionPlanByID)
	subscriptionPlan.Get("/user/:userId", jwtHandler.ValidateToken, subscriptionHandler.GetSubscriptionPlanByUserId)
	subscriptionPlan.Put("/:subscriptionPlanId", jwtHandler.ValidateToken, subscriptionHandler.UpdateSubscriptionPlanByID)
	subscriptionPlan.Delete("/:subscriptionPlanId", jwtHandler.ValidateToken, subscriptionHandler.DeleteSubscriptionPlanByID)
}

func SetupSubscriptionRelatedRoutes(router fiber.Router) {
	SetupSubscriptionRoutes(router)
	SetupSubscriptionPlanRoutes(router)
}
