package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateRecord(ctx *fiber.Ctx, model interface{}) error {
	modelBody := model
	modelName := strings.ToLower(reflect.TypeOf(model).Elem().Name())
	err := ctx.BodyParser(modelBody)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "request failed",
			"data":    err})
		return err
	}

	err = storage.DB.Create(modelBody).Error
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create " + modelName,
			"data":    err})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create " + modelName + " succeeded",
		"data":    modelBody})
	return nil
}

func GetRecords(ctx *fiber.Ctx, models interface{}) error {
	modelName := strings.ToLower(reflect.TypeOf(models).Elem().Elem().Name()) + "s"
	err := storage.DB.Find(models).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch " + modelName,
			"data":    err})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": modelName + " fetched successfully",
		"data":    models})
	return nil
}

func GetRecordByID(ctx *fiber.Ctx, model interface{}, idParamName string) error {
	modelName := strings.ToLower(reflect.TypeOf(model).Elem().Name())
	id := ctx.Params(idParamName)
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": modelName + " ID cannot be empty on get",
			"data":    nil})
	}
	res := storage.DB.Debug().Where("id = ?", id).Find(model)
	if res.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "error while fetching the " + modelName,
			"data":    res.Error})
		return res.Error
	} else if res.RowsAffected == 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "could not find the " + modelName + " with id " + id})
		return nil
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": modelName + " found successfully",
		"data":    model})
	return nil
}

func DeleteRecordByID(ctx *fiber.Ctx, model interface{}, idParamName string) error {
	modelName := strings.ToLower(reflect.TypeOf(model).Elem().Name())
	id := ctx.Params(idParamName)
	if id == "" {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": modelName + " ID cannot be empty on delete"})
	}

	switch modelName {
	case "user":
		_, err := uuid.Parse(id)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"message": "invalid ID for " + modelName,
				"data":    err.Error(),
			})
		}
	}

	res := storage.DB.Where("id = ?", id).Unscoped().Delete(&model)
	if res.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "error while deleting " + modelName,
			"data":    res.Error})
		return res.Error
	} else if res.RowsAffected == 0 {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "couldn't find " + modelName + " with id " + id})
		return nil
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": modelName + " deleted successfully"})
	return nil
}

func UpdateRecordByID(ctx *fiber.Ctx, model interface{}, updatingStruct interface{}, idParamName string) error {
	modelName := strings.ToLower(reflect.TypeOf(model).Elem().Name())
	err := GetRecordByID(ctx, model, idParamName)
	if err != nil {
		return err
	}
	toBeUpdatedStructType := reflect.TypeOf(model).Elem()
	toBeUpdatedStructInstance := reflect.New(toBeUpdatedStructType)

	updatingFieldsStructValue := reflect.ValueOf(updatingStruct).Elem()
	typeOfUpdatingFieldsStructValue := updatingFieldsStructValue.Type()

	var permittedUpdateFields []string
	for i := 0; i < updatingFieldsStructValue.NumField(); i++ {
		permittedUpdateFields = append(permittedUpdateFields, typeOfUpdatingFieldsStructValue.Field(i).Name)
	}

	err = ctx.BodyParser(updatingStruct)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			// how do I retrieve and display all field names of fieldsToBeUpdatedStruct?
			"message": "bad input: you can only update " + strings.Join(permittedUpdateFields, ", ") + " fields",
			"data":    err})
		return err
	}

	for i := 0; i < updatingFieldsStructValue.NumField(); i++ {
		fieldValue := updatingFieldsStructValue.Field(i)
		fieldName := typeOfUpdatingFieldsStructValue.Field(i).Name

		toBeUpdatedField := toBeUpdatedStructInstance.Elem().FieldByName(fieldName)
		if toBeUpdatedField.IsValid() && toBeUpdatedField.CanSet() {
			toBeUpdatedField.Set(fieldValue)
		}
	}

	storage.DB.Model(model).Updates(toBeUpdatedStructInstance.Interface())

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": modelName + " updated successfully",
		"data":    model})
	return nil
}
