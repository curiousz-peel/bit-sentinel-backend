package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	Type     string  `json:"type" gorm:"unique"`
	Duration int     `json:"duration"`
	Price    float32 `json:"price"`
}

type SubscriptionPlan struct {
	gorm.Model
	UserID         uuid.UUID    `json:"userId" gorm:"uniqueIndex:idx_user_subscription;not null;default:null"`
	User           User         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SubscriptionID uint         `json:"subscriptionId" gorm:"uniqueIndex:idx_user_subscription;not null;default:null"`
	Subscription   Subscription `gorm:"foreignKey:SubscriptionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StartDate      time.Time    `json:"startDate" gorm:"default:null"`
	EndDate        time.Time    `json:"endDate" gorm:"default:null"`
}

type UpdateSubscriptionPlan struct {
	UserID         uuid.UUID `json:"userID"`
	SubscriptionID int       `json:"subscriptionID"`
}

type SubscriptionPlanDTO struct {
	ID               uint      `json:"id"`
	UserID           uuid.UUID `json:"userId"`
	SubscriptionID   uint      `json:"subscriptionId"`
	SubscriptionType string    `json:"subscriptionType"`
	StartDate        time.Time `json:"startDate"`
	EndDate          time.Time `json:"endDate"`
}

func ToSubscriptionPlanDTO(subscriptionPlan SubscriptionPlan) SubscriptionPlanDTO {
	return SubscriptionPlanDTO{
		ID:               subscriptionPlan.ID,
		UserID:           subscriptionPlan.UserID,
		SubscriptionID:   subscriptionPlan.SubscriptionID,
		StartDate:        subscriptionPlan.StartDate,
		EndDate:          subscriptionPlan.EndDate,
		SubscriptionType: subscriptionPlan.Subscription.Type,
	}
}

type UpdateSubscription struct {
	Type     string  `json:"type"`
	Duration int     `json:"duration"`
	Price    float32 `json:"price"`
}
