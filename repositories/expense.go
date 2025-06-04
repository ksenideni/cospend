package repositories

import (
	"cospend/models"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	DB *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{DB: db}
}

func (r *ExpenseRepository) CreateExpense(expense *models.Expense) error {
	return r.DB.Create(expense).Error
}

func (r *ExpenseRepository) GetExpensesByGroupID(groupID int) ([]models.Expense, error) {
	var expenses []models.Expense
	err := r.DB.Where("group_id = ?", groupID).Find(&expenses).Error
	return expenses, err
}

func (r *ExpenseRepository) GetExpenseByID(expenseID int) (models.Expense, error) {
	var expense models.Expense
	err := r.DB.First(&expense, expenseID).Error
	return expense, err
}
