package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetSubscriptionPlans() (*[]models.SubscriptionPlanDTO, error) {
	subscriptionPlans := &[]models.SubscriptionPlan{}
	subscriptionPlansDTOs := []models.SubscriptionPlanDTO{}
	err := storage.DB.Model(&models.SubscriptionPlan{}).Preload("User").Preload("Subscription").Find(&subscriptionPlans).Error
	if err != nil {
		return nil, err
	}
	for _, subscriptionPlan := range *subscriptionPlans {
		subscriptionPlansDTOs = append(subscriptionPlansDTOs, models.ToSubscriptionPlanDTO(subscriptionPlan))
	}
	return &subscriptionPlansDTOs, nil
}

func CreateSubscriptionPlan(subscriptionPlan *models.SubscriptionPlan) (*models.SubscriptionPlanDTO, error) {
	res := storage.DB.Debug().Find(&subscriptionPlan.User, "id = ?", subscriptionPlan.UserID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding user: " + subscriptionPlan.UserID.String())
	}
	res = storage.DB.Find(&subscriptionPlan.Subscription, "id = ?", subscriptionPlan.SubscriptionID)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("error in finding subscription type: " + subscriptionPlan.Subscription.Type + " with error " + res.Error.Error())
	}
	subscriptionPlan.StartDate = time.Now()
	if subscriptionPlan.SubscriptionID != 1 {
		subscriptionPlan.EndDate = time.Now().Add(time.Hour * 24 * time.Duration(subscriptionPlan.Subscription.Duration))
	}
	err := storage.DB.Omit("User").Create(subscriptionPlan).Error
	if err != nil {
		return nil, err
	}
	subscriptionPlanDTO := models.ToSubscriptionPlanDTO(*subscriptionPlan)
	return &subscriptionPlanDTO, nil
}

func GetSubscriptionPlanByID(id string) (*models.SubscriptionPlanDTO, error) {
	subscriptionPlan := &models.SubscriptionPlan{}
	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Subscription").Find(subscriptionPlan)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("could not find subscription plan with id " + id)
	}
	subscriptionPlanDTO := models.ToSubscriptionPlanDTO(*subscriptionPlan)
	return &subscriptionPlanDTO, nil
}

func GetSubscriptionPlanByUserId(id string) (*models.SubscriptionPlanDTO, error) {
	subscriptionPlan := &models.SubscriptionPlan{}
	res := storage.DB.Where("user_id = ?", id).Preload("User").Preload("Subscription").Find(subscriptionPlan)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("could not find subscription plan with user id " + id)
	}
	subscriptionPlanDTO := models.ToSubscriptionPlanDTO(*subscriptionPlan)
	return &subscriptionPlanDTO, nil
}

func DeleteSubscriptionPlanByID(id string) error {
	subscriptionPlan := &models.SubscriptionPlan{}
	res := storage.DB.Unscoped().Delete(subscriptionPlan, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete subscription plan with id " + id)
	}
	return nil
}

func UpdateSubscriptionPlanByID(id string, updateSubscriptionPlan models.UpdateSubscriptionPlan) error {
	subscriptionPlan := &models.SubscriptionPlan{}
	res := storage.DB.Where("id = ?", id).Preload("User").Preload("Subscription").Find(subscriptionPlan)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the subscription plan with id " + id)
	}
	user := &models.User{}
	res = storage.DB.Where("id = ?", updateSubscriptionPlan.UserID).Find(user)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the subscription plan user with id " + updateSubscriptionPlan.UserID.String())
	}
	subscription := &models.Subscription{}
	res = storage.DB.Where("id = ?", updateSubscriptionPlan.SubscriptionID).Find(subscription)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the subscription plan type with id " + fmt.Sprint(updateSubscriptionPlan.SubscriptionID))
	}
	storage.DB.Model(&subscriptionPlan).Updates(&models.SubscriptionPlan{
		UserID:         user.ID,
		User:           *user,
		SubscriptionID: subscription.ID,
		Subscription:   *subscription})
	return nil
}
