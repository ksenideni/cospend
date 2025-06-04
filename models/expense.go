package models

import "time"

type Expense struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID     int       `json:"groupId" gorm:"not null;index"`
	CreatedBy   string    `json:"createdBy,omitempty" gorm:"index"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Amount      float64   `json:"amount" gorm:"type:numeric(10,2);not null"`
	Date        time.Time `json:"date" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"default:now()"`
}

func (Expense) TableName() string {
	return "expenses"
}

// DTOs

type ExpenseRequest struct {
	Description string    `json:"description" binding:"required"`
	Amount      float64   `json:"amount" binding:"required,gt=0"`
	Date        time.Time `json:"date" binding:"required"`
}

type ListExpense struct {
	Page      int       `json:"page"`
	Limit     int       `json:"limit"`
	Total     int       `json:"total"`
	TotalPage int       `json:"totalPage"`
	Expenses  []Expense `json:"expenses"`
}
