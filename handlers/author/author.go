package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAuthors(ctx *fiber.Ctx) error {
	authors := &[]models.Author{}
	err := storage.DB.Model(&models.Author{}).Preload("User").Find(&authors).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch authors",
			"data":    err})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "authors fetched successfully",
		"data":    authors})
	return nil
}

func CreateAuthor(ctx *fiber.Ctx) error {
	author := &models.Author{}
	err := ctx.BodyParser(&author)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "request failed",
			"data":    err})
		return err
	}

	res := storage.DB.Find(&author.User, "id = ?", author.UserID)
	if res.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "error while fetching the user for a new author",
			"data":    res.Error})
		return res.Error
	} else if res.RowsAffected == 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "could not find a user with id: " + author.UserID.String() + " to link to a new author"})
		return nil
	}

	err = storage.DB.Debug().Omit("User").Create(author).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create author",
			"data":    err})
		fmt.Println(err)
		return err
	}

	type updateUserAuthorStatus struct {
		IsAuthor bool `json:"isAuthor"`
	}

	res = storage.DB.Model(&author.User).Updates(updateUserAuthorStatus{IsAuthor: true})
	if res.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not update user author status",
			"data":    res.Error})
		fmt.Println(res.Error)
		return res.Error
	} else if res.RowsAffected == 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "could not update the user author status to true"})
		return nil
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create author succeeded",
		"data":    author})
	return nil

}

func GetAuthorByID(ctx *fiber.Ctx) error {
	author := &models.Author{}
	id := ctx.Params("authorId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "author ID cannot be empty on get",
			"data":    nil})
	}

	res := storage.DB.Where("id = ?", id).Preload("User").Find(author)
	if res.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the author",
			"data":    res.Error})
		return res.Error
	} else if res.RowsAffected == 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "could not find author with id: " + id + " to fetch"})
		return errors.New("could not find author with id: " + id + " to fetch")
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "author found successfully",
		"data":    author})
	return nil
}

func DeleteAuthorByID(ctx *fiber.Ctx) error {
	author := &models.Author{}
	id := ctx.Params("authorId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "author ID cannot be empty on delete",
			"data":    nil})
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid author ID",
			"data":    err.Error(),
		})
	}

	res := storage.DB.Unscoped().Delete(author, uuid)
	if res.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete author",
			"data":    err})
		return err
	} else if res.RowsAffected == 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "could not find a user with id: " + uuid.String() + " to delete"})
		return nil
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "author deleted successfully"})
	return nil
}

func UpdateAuthorByID(ctx *fiber.Ctx) error {
	type updateAuthor struct {
		Profession  string `json:"profession"`
		Description string `json:"description"`
		Topics      string `json:"topics"`
	}
	author := &models.Author{}
	id := ctx.Params("authorId")
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "author ID cannot be empty on update",
			"data":    nil})
	}

	res := storage.DB.Where("id = ?", id).Preload("User").Find(author)
	if res.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not find the author",
			"data":    res.Error})
		return res.Error
	} else if res.RowsAffected == 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "could not find the author"})
		return nil
	}

	var updateAuthorData updateAuthor
	err := ctx.BodyParser(&updateAuthorData)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "bad input: you can only update profession, description or topics fields",
			"data":    err})
		return err
	}

	storage.DB.Model(&author).Updates(&models.Author{
		Profession:  updateAuthorData.Profession,
		Description: updateAuthorData.Description,
		Topics:      updateAuthorData.Topics})

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "author updated successfully",
		"data":    author})
	return nil
}
