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
	UserID         uuid.UUID    `json:"userId" gorm:"not null;unique"`
	User           User         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SubscriptionID uint         `json:"subscriptionId" gorm:"not null;default:null"`
	Subscription   Subscription `gorm:"foreignKey:SubscriptionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StartDate      time.Time    `json:"startDate" gorm:"default:null"`
	EndDate        time.Time    `json:"endDate" gorm:"default:null"`
}

// func (p *SubscriptionPlan) AfterCreate(tx *gorm.DB) (err error) {
// 	res := storage.DB.Find(&p.Subscription, "id = ?", p.SubscriptionID)
// 	if res.Error != nil {
// 		return errors.New("error in finding the subscription with id " + fmt.Sprint(p.SubscriptionID) + " with error: " + res.Error.Error())
// 	} else if res.RowsAffected == 0 {
// 		return errors.New("could not find the subscription with id " + fmt.Sprint(p.SubscriptionID) + " with error: " + res.Error.Error())
// 	}
// 	storage.DB.Model(&p).Updates(map[string]interface{}{"StartDate": time.Now(),
// 		"EndDate": time.Now().Add(time.Hour * 24 * time.Duration(p.Subscription.Duration))})
// 	return nil
// }
