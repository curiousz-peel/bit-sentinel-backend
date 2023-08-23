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

type UpdateSubscription struct {
	Type     string  `json:"type"`
	Duration int     `json:"duration"`
	Price    float32 `json:"price"`
}
