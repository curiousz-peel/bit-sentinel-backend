package routes

import (
	mediaHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/media"
	"github.com/gofiber/fiber/v2"
)

func SetupMediaRoutes(router fiber.Router) {
	media := router.Group("/media")
	media.Post("/", mediaHandler.CreateMedia)
	media.Get("/", mediaHandler.GetMedias)
	media.Get("/:mediaId", mediaHandler.GetMediaByID)
	media.Put("/:mediaId", mediaHandler.UpdateMediaByID)
	media.Delete("/:mediaId", mediaHandler.DeleteMediaByID)
}

func SetupMediaTypeRoutes(router fiber.Router) {
	mediaType := router.Group("/mediaType")
	mediaType.Post("/", mediaHandler.CreateMediaType)
	mediaType.Get("/", mediaHandler.GetMediaTypes)
	mediaType.Get("/:mediaTypeId", mediaHandler.GetMediaTypeByID)
	mediaType.Get("/:mediaTypeName", mediaHandler.GetMediaTypeByType)
	mediaType.Put("/:mediaTypeId", mediaHandler.UpdateMediaTypeByID)
	mediaType.Delete("/:mediaTypeId", mediaHandler.DeleteMediaTypeByID)
}

func SetupMediaRelatedRoutes(router fiber.Router) {
	SetupMediaRoutes(router)
	SetupMediaTypeRoutes(router)
}
