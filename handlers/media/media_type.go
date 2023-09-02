package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/media"
	"github.com/gofiber/fiber/v2"
)

func GetMediaTypes(ctx *fiber.Ctx) error {
	mediaTypes, err := service.GetMediaTypes()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch media types",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "media types fetched successfully",
		"data":    mediaTypes})
}

func CreateMediaType(ctx *fiber.Ctx) error {
	mediaType := &models.MediaType{}
	err := ctx.BodyParser(mediaType)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err})
		return err
	}
	mediaType, err = service.CreateMediaType(mediaType)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create media type",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "media type created successfully",
		"data":    mediaType})
}

func GetMediaTypeByID(ctx *fiber.Ctx) error {
	id := ctx.Params("mediaTypeId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "media type ID cannot be empty on get",
			"data":    nil})
	}
	mediaType, err := service.GetMediaTypeByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find media type, check if ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "media type fetched successfully",
		"data":    mediaType})
}

func GetMediaTypeByType(ctx *fiber.Ctx) error {
	typeName := ctx.Params("mediaTypeName")
	if typeName == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "media type name cannot be empty on get",
			"data":    nil})
	}
	mediaType, err := service.GetMediaTypeByType(typeName)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find media type, check if name " + typeName + " exists",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "media type fetched successfully",
		"data":    mediaType})
}

func DeleteMediaTypeByID(ctx *fiber.Ctx) error {
	id := ctx.Params("mediaTypeId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "media type ID cannot be empty on delete",
			"data":    nil})
	}
	err := service.DeleteMediaTypeByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete media type, check if media with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "media type deleted successfully"})
}

func UpdateMediaTypeByID(ctx *fiber.Ctx) error {
	id := ctx.Params("mediaTypeId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "media type ID cannot be empty on update",
			"data":    nil})
	}
	var updateMediaTypeData models.UpdateMediaType
	err := ctx.BodyParser(&updateMediaTypeData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the type string",
			"data":    err})
	}
	err = service.UpdateMediaTypeByID(id, updateMediaTypeData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update media type",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "media type updated successfully",
		"data":    updateMediaTypeData})
}
