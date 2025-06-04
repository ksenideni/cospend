package services

import (
	"cospend/models"
	"cospend/repositories"
)

type ExpenseService struct {
	ExpenseRepo repositories.ExpenseRepository
}

func NewExpenseService(repo repositories.ExpenseRepository) *ExpenseService {
	return &ExpenseService{ExpenseRepo: repo}
}

func (s *ExpenseService) CreateExpense(expense *models.Expense) error {
	return s.ExpenseRepo.CreateExpense(expense)
}

func (s *ExpenseService) GetExpensesByGroupID(groupID int) ([]models.Expense, error) {
	return s.ExpenseRepo.GetExpensesByGroupID(groupID)
}

func (s *ExpenseService) GetExpenseByID(expenseID int) (models.Expense, error) {
	return s.ExpenseRepo.GetExpenseByID(expenseID)
}
