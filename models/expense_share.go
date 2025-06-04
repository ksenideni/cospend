package models

type ExpenseShare struct {
	ID        int     `json:"id" gorm:"primaryKey;autoIncrement"`
	ExpenseID int     `json:"expenseId" gorm:"not null;index"`
	UserID    string  `json:"userId" gorm:"not null;index"`
	Paid      float64 `json:"paid" gorm:"type:numeric(10,2);default:0"`
	Owed      float64 `json:"owed" gorm:"type:numeric(10,2);default:0"`
}

func (ExpenseShare) TableName() string {
	return "expense_shares"
}

type ExpenseShareInput struct {
	UserID string  `json:"userId" binding:"required"`
	Paid   float64 `json:"paid"`
	Owed   float64 `json:"owed"`
}

type ExpenseSharesResponse struct {
	ExpenseID int             `json:"expenseId"`
	Shares    []ExpenseShare  `json:"shares"`
}
