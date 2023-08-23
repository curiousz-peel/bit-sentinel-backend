package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	subscriptionPlanService "github.com/curiousz-peel/web-learning-platform-backend/service/subscription"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func CreateUser(user *models.User) (*models.User, error) {
	err := storage.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	_, err = subscriptionPlanService.CreateSubscriptionPlan(&models.SubscriptionPlan{UserID: user.ID, SubscriptionID: 1, StartDate: time.Now(), EndDate: time.Now().AddDate(1000, 0, 0)})
	if err != nil {
		return nil, errors.New("failed to create basic subscription plan for new user with id " + user.ID.String())
	}
	return user, nil
}

func UpdateUserByID(id string, updateUser models.UpdateUser) error {
	user := &models.User{}
	res := storage.DB.Where("id = ?", id).Find(user)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the user, check if ID " + id + " exists")
	}
	storage.DB.Model(&user).Updates(&models.User{
		FirstName: updateUser.FirstName,
		LastName:  updateUser.LastName,
		UserName:  updateUser.UserName,
		Email:     updateUser.Email,
		Password:  updateUser.Password,
		Birthday:  updateUser.Birthday,
		IsMod:     updateUser.IsMod})
	return nil
}

func GetUsers() (*[]models.User, error) {
	users := []models.User{}
	err := storage.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		storage.DB.Where("user_id = ?", user.ID).Find(&user.Ratings)
		storage.DB.Where("user_id = ?", user.ID).Find(&user.Enrollments)
		storage.DB.Where("user_id = ?", user.ID).Find(&user.Comments)
		fmt.Println(user)
	}
	return &users, nil
}

func GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	res := storage.DB.Where("id = ?", id).Find(user)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("can't find user with id " + id)
	}
	storage.DB.Where("user_id = ?", user.ID).Find(&user.Ratings)
	storage.DB.Where("user_id = ?", user.ID).Find(&user.Enrollments)
	storage.DB.Where("user_id = ?", user.ID).Find(&user.Comments)
	return user, nil
}

func DeleteUserByID(id string) error {
	user := &models.User{}
	res := storage.DB.Unscoped().Delete(user, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete user with id " + id)
	}
	return nil
}
