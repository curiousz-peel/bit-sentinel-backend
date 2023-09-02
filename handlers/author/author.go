package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/author"

	"github.com/gofiber/fiber/v2"
)

func GetAuthors(ctx *fiber.Ctx) error {
	authors, err := service.GetAuthors()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch authors",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "authors fetched successfully",
		"data":    authors})
}

func GetAuthorByID(ctx *fiber.Ctx) error {
	id := ctx.Params("authorId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "author ID cannot be empty on get",
			"data":    nil})
	}
	author, err := service.GetAuthorByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the author, check if ID " + id + " exists",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "author fetched successfully",
		"data":    author})
}

func DeleteAuthorByID(ctx *fiber.Ctx) error {
	id := ctx.Params("authorId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "author ID cannot be empty on delete",
			"data":    nil})
	}

	err := service.DeleteAuthorByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete author, check if author with ID " + id + " exists",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "author deleted successfully"})
}

func CreateAuthor(ctx *fiber.Ctx) error {
	author := &models.Author{}
	err := ctx.BodyParser(author)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err})
		return err
	}

	authorResponse, err := service.CreateAuthor(author)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create author",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "author created successfully",
		"data":    authorResponse})
}

func UpdateAuthorByID(ctx *fiber.Ctx) error {
	id := ctx.Params("authorId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "author ID cannot be empty on update",
			"data":    nil})
	}

	var updateAuthorData models.UpdateAuthor
	err := ctx.BodyParser(&updateAuthorData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the user id, the profession, the description or the topics",
			"data":    err})
	}

	err = service.UpdateAuthorByID(id, updateAuthorData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the author",
			"data":    err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "author updated successfully",
		"data":    updateAuthorData})
}
