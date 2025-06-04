package models

type Debt struct {
	ID         int     `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID    int     `json:"groupId" gorm:"not null;index"`
	FromUserID string  `json:"fromUserId" gorm:"not null;index"`
	ToUserID   string  `json:"toUserId" gorm:"not null;index"`
	Amount     float64 `json:"amount" gorm:"type:numeric(10,2);not null"`
}

func (Debt) TableName() string {
	return "debts"
}
// DTOs

type DebtDto struct {
	ToUserID   string  `json:"to_user_id"`
	ToUserName string  `json:"to_user_name"`
	Amount     float64 `json:"amount"`
}
