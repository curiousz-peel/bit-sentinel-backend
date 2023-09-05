package service

import (
	"errors"
	"time"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	subscriptionPlanService "github.com/curiousz-peel/web-learning-platform-backend/service/subscription"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/google/uuid"
)

func CreateUser(user *models.User) (*models.UserDTO, error) {
	err := storage.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	_, err = subscriptionPlanService.CreateSubscriptionPlan(&models.SubscriptionPlan{UserID: user.ID, SubscriptionID: 1, StartDate: time.Now(), EndDate: time.Now().AddDate(1000, 0, 0)})
	if err != nil {
		return nil, errors.New("failed to create basic subscription plan for new user with id " + user.ID.String())
	}
	userDTO := models.ToUserDTO(*user)
	return &userDTO, nil
}

func UpdateUserByID(id string, updateUser models.UpdateUser) error {
	user := &models.User{}
	res := storage.DB.Where("id = ?", id).Find(user)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the user, check if ID " + id + " exists")
	}
	if !updateUser.IsAuthor {
		storage.DB.Debug().Model(&user).Updates(map[string]interface{}{
			"IsAuthor": false,
		})
	}
	if !updateUser.IsMod {
		storage.DB.Debug().Model(&user).Updates(map[string]interface{}{
			"IsMod": false,
		})
	}
	storage.DB.Debug().Model(&user).Updates(&models.User{
		FirstName: updateUser.FirstName,
		LastName:  updateUser.LastName,
		UserName:  updateUser.UserName,
		Email:     updateUser.Email,
		Password:  updateUser.Password,
		Birthday:  updateUser.Birthday,
		IsMod:     updateUser.IsMod,
		IsAuthor:  updateUser.IsAuthor})
	return nil
}

func GetUsers() (*[]models.UserDTO, error) {
	users := []models.User{}
	usersDTOs := []models.UserDTO{}
	err := storage.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	for i := range users {
		storage.DB.Where("user_id = ?", users[i].ID).Find(&users[i].Ratings)
		storage.DB.Where("user_id = ?", users[i].ID).Find(&users[i].Enrollments)
		storage.DB.Where("user_id = ?", users[i].ID).Find(&users[i].Comments)

		usersDTOs = append(usersDTOs, models.ToUserDTO(users[i]))
	}
	return &usersDTOs, nil
}

func GetUserByID(id string) (*models.UserDTO, error) {
	user := &models.User{}
	res := storage.DB.Where("id = ?", id).Find(user)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("can't find user with id " + id)
	}
	storage.DB.Where("user_id = ?", user.ID).Find(&user.Ratings)
	storage.DB.Where("user_id = ?", user.ID).Find(&user.Enrollments)
	storage.DB.Where("user_id = ?", user.ID).Find(&user.Comments)
	userDTO := models.ToUserDTO(*user)
	return &userDTO, nil
}

func GetUserByUsername(userName string) (*models.UserDTO, error) {
	user := &models.User{}
	res := storage.DB.Where("user_name = ?", userName).Find(user)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("can't find user with username " + userName)
	}
	storage.DB.Where("user_id = ?", user.ID).Find(&user.Ratings)
	storage.DB.Where("user_id = ?", user.ID).Find(&user.Enrollments)
	storage.DB.Where("user_id = ?", user.ID).Find(&user.Comments)
	userDTO := models.ToUserDTO(*user)
	return &userDTO, nil
}

func DeleteUserByID(id string) error {
	user := &models.User{}
	userUUID, _ := uuid.Parse(id)
	res := storage.DB.Unscoped().Delete(user, userUUID)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete user with id " + id)
	}
	return nil
}
