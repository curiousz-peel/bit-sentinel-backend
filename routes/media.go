package routes

import (
	mediaHandler "github.com/curiousz-peel/web-learning-platform-backend/handlers/media"
	jwtHandler "github.com/curiousz-peel/web-learning-platform-backend/requestValidator"
	"github.com/gofiber/fiber/v2"
)

func SetupMediaRoutes(router fiber.Router) {
	media := router.Group("/media")
	media.Post("/", jwtHandler.ValidateToken, mediaHandler.CreateMedia)
	media.Get("/", jwtHandler.ValidateToken, mediaHandler.GetMedias)
	media.Get("/:mediaId", jwtHandler.ValidateToken, mediaHandler.GetMediaByID)
	media.Put("/:mediaId", jwtHandler.ValidateToken, mediaHandler.UpdateMediaByID)
	media.Delete("/:mediaId", jwtHandler.ValidateToken, mediaHandler.DeleteMediaByID)
}

func SetupMediaTypeRoutes(router fiber.Router) {
	mediaType := router.Group("/mediaType")
	mediaType.Post("/", jwtHandler.ValidateToken, mediaHandler.CreateMediaType)
	mediaType.Get("/", jwtHandler.ValidateToken, mediaHandler.GetMediaTypes)
	mediaType.Get("/:mediaTypeId", jwtHandler.ValidateToken, mediaHandler.GetMediaTypeByID)
	mediaType.Get("/:mediaTypeName", jwtHandler.ValidateToken, mediaHandler.GetMediaTypeByType)
	mediaType.Put("/:mediaTypeId", jwtHandler.ValidateToken, mediaHandler.UpdateMediaTypeByID)
	mediaType.Delete("/:mediaTypeId", jwtHandler.ValidateToken, mediaHandler.DeleteMediaTypeByID)
}

func SetupMediaRelatedRoutes(router fiber.Router) {
	SetupMediaRoutes(router)
	SetupMediaTypeRoutes(router)
}
