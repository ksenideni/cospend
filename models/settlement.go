package models

import "time"

type Settlement struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID    int       `json:"groupId" gorm:"not null;index"`
	FromUserID string    `json:"fromUserId" gorm:"not null;index"`
	ToUserID   string    `json:"toUserId" gorm:"not null;index"`
	Amount     float64   `json:"amount" gorm:"type:numeric(10,2);not null"`
	SettledAt  time.Time `json:"settledAt" gorm:"default:now()"`
}

func (Settlement) TableName() string {
	return "settlements"
}



type SettleRequest struct {
	ToUserID string  `json:"to_user_id" binding:"required"`
	Amount   float64 `json:"amount" binding:"required,gt=0"`
}
