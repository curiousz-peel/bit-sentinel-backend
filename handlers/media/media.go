package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/media"
	"github.com/gofiber/fiber/v2"
)

func GetMedias(ctx *fiber.Ctx) error {
	medias, err := service.GetMedias()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch medias",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "medias fetched successfully",
		"data":    medias})
}

func GetMediaByID(ctx *fiber.Ctx) error {
	id := ctx.Params("mediaId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "media ID cannot be empty on get",
			"data":    nil})
	}
	media, err := service.GetMediaByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find media, check if ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "media fetched successfully",
		"data":    media})
}

func DeleteMediaByID(ctx *fiber.Ctx) error {
	id := ctx.Params("mediaId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "media ID cannot be empty on delete",
			"data":    nil})
	}
	err := service.DeleteMediaByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete media, check if media with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "media deleted successfully"})
}

func CreateMedia(ctx *fiber.Ctx) error {
	media := &models.Media{}
	err := ctx.BodyParser(media)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err})
		return err
	}
	mediaDTO, err := service.CreateMedia(media)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create media",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "media created successfully",
		"data":    mediaDTO})
}

func UpdateMediaByID(ctx *fiber.Ctx) error {
	id := ctx.Params("mediaId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "media ID cannot be empty on update",
			"data":    nil})
	}
	var updateMediaData models.UpdateMedia
	err := ctx.BodyParser(&updateMediaData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the lesson id, file path or file type name",
			"data":    err})
	}
	err = service.UpdateMediaByID(id, updateMediaData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update media",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "media updated successfully",
		"data":    updateMediaData})
}
