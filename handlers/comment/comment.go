package handlers

import (
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	service "github.com/curiousz-peel/web-learning-platform-backend/service/comment"

	"github.com/gofiber/fiber/v2"
)

func GetComments(ctx *fiber.Ctx) error {
	comments, err := service.GetComments()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch comments",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "comments fetched successfully",
		"data":    comments})
}

func GetCommentByID(ctx *fiber.Ctx) error {
	id := ctx.Params("commentId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "comment ID cannot be empty on get",
			"data":    nil})
	}
	comment, err := service.GetCommentByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the comment, check if ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "comment fetched successfully",
		"data":    comment})
}

func DeleteCommentByID(ctx *fiber.Ctx) error {
	id := ctx.Params("commentId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "comment ID cannot be empty on delete",
			"data":    nil})
	}

	err := service.DeleteCommentByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete comment, check if comment with ID " + id + " exists",
			"data":    err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "comment deleted successfully"})
}

func CreateComment(ctx *fiber.Ctx) error {
	comment := &models.Comment{}
	err := ctx.BodyParser(comment)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "failed to parse request body",
			"data":    err.Error()})
		return err
	}

	commentResponse, err := service.CreateComment(comment)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create comment",
			"data":    err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "comment created successfully",
		"data":    commentResponse})
}

func UpdateCommentByID(ctx *fiber.Ctx) error {
	id := ctx.Params("commentId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "comment ID cannot be empty on update",
			"data":    nil})
	}

	var updateCommentData models.UpdateComment
	err := ctx.BodyParser(&updateCommentData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update the user id, the course id or the text",
			"data":    err})
	}

	err = service.UpdateCommentByID(id, updateCommentData)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the comment",
			"data":    err})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "comment updated successfully",
		"data":    updateCommentData})
}
