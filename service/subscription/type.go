package service

import (
	"errors"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
)

func GetSubscriptions() (*[]models.Subscription, error) {
	subscriptions := &[]models.Subscription{}
	err := storage.DB.Model(&models.Subscription{}).Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func GetSubscriptionByID(id string) (*models.Subscription, error) {
	subscription := &models.Subscription{}
	res := storage.DB.Where("id = ?", id).Find(subscription)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New("could not find subscription with id " + id)
	}
	return subscription, nil
}

func DeleteSubscriptionByID(id string) error {
	subscription := &models.Subscription{}
	res := storage.DB.Unscoped().Delete(subscription, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not delete subscription with id " + id)
	}
	return nil
}

func CreateSubscription(subscription *models.Subscription) (*models.Subscription, error) {
	err := storage.DB.Debug().Create(subscription).Error
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func UpdateSubscriptionByID(id string, updateSubscription models.UpdateSubscription) error {
	subscription := &models.Subscription{}
	res := storage.DB.Where("id = ?", id).Find(subscription)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("could not find the subscription with id " + id)
	}

	storage.DB.Model(&subscription).Updates(&models.Subscription{
		Type:     updateSubscription.Type,
		Duration: updateSubscription.Duration,
		Price:    updateSubscription.Price})
	return nil
}
