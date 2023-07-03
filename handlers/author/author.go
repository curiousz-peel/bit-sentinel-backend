package handlers

import (
	"fmt"
	"net/http"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/gofiber/fiber/v2"
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

	fmt.Println("here")

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
	fmt.Println(author.UserID)
	fmt.Println(author.User)

	err = storage.DB.Debug().Omit("User").Create(author).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create author",
			"data":    err})
		fmt.Println(err)
		return err
	}

	// ???
	// storage.DB.Debug().Model(&author).Association("User").Replace(&author.User)
	err = storage.DB.Model(author).Updates(models.Author{User: models.User{}}).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create author2",
			"data":    err})
		fmt.Println(err)
		return err
	}

	fmt.Println("why god")

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create author succeeded",
		"data":    author})
	return nil

}

// func CreateUser(ctx *fiber.Ctx) error {
// 	user := models.User{}
// 	err := ctx.BodyParser(&user)
// 	if err != nil {
// 		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
// 			"message": "request failed",
// 			"data":    err})
// 		return err
// 	}

// 	err = storage.DB.Create(&user).Error
// 	if err != nil {
// 		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
// 			"message": "could not create user",
// 			"data":    err})
// 		return err
// 	}

// 	ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "create user succeeded",
// 		"data":    user})
// 	return nil
// }

// func GetUserByID(ctx *fiber.Ctx) error {
// 	user := &models.User{}
// 	id := ctx.Params("userId")
// 	if id == "" {
// 		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
// 			"message": "user ID cannot be empty on get",
// 			"data":    nil})
// 	}
// 	fmt.Println("the user id is", id)
// 	err := storage.DB.Where("id = ?", id).Find(user).Error
// 	if err != nil {
// 		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
// 			"message": "could not find the user",
// 			"data":    err})
// 		return err
// 	}
// 	ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "user found successfully",
// 		"data":    user})
// 	return nil
// }

// func DeleteUserByID(ctx *fiber.Ctx) error {
// 	user := &models.User{}
// 	id := ctx.Params("userId")
// 	if id == "" {
// 		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
// 			"message": "user ID cannot be empty on delete"})
// 	}

// 	uuid, err := uuid.Parse(id)
// 	if err != nil {
// 		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
// 			"message": "invalid user ID",
// 			"data":    err.Error(),
// 		})
// 	}

// 	err = storage.DB.Delete(user, uuid).Error
// 	if err != nil {
// 		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
// 			"message": "could not delete user",
// 			"data":    err})
// 		return err
// 	}
// 	ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "user deleted successfully"})
// 	return nil
// }

// func UpdateUserByID(ctx *fiber.Ctx) error {
// 	type updateUser struct {
// 		FirstName string `json:"firstName"`
// 		LastName  string `json:"lastName"`
// 		Email     string `json:"email"`
// 		Password  string `json:"password"`
// 	}
// 	user := &models.User{}
// 	id := ctx.Params("userId")
// 	if id == "" {
// 		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
// 			"message": "user ID cannot be empty on get",
// 			"data":    nil})
// 	}
// 	fmt.Println("the user id is", id)
// 	err := storage.DB.Where("id = ?", id).Find(user).Error
// 	if err != nil {
// 		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
// 			"message": "could not find the user",
// 			"data":    err})
// 		return err
// 	}

// 	var updateUserData updateUser
// 	err = ctx.BodyParser(&updateUserData)
// 	if err != nil {
// 		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"message": "bad input: you can only update first name, last name, email or password fields",
// 			"data":    err})
// 		return err
// 	}

// 	storage.DB.Model(&user).Updates(&models.User{
// 		FirstName: updateUserData.FirstName,
// 		LastName:  updateUserData.LastName,
// 		Email:     updateUserData.Email,
// 		Password:  updateUserData.Password})

// 	ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "user updated successfully",
// 		"data":    user})
// 	return nil
// }
