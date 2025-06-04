package services

import (
	"cospend/models"
	"cospend/repositories"
)

type DebtService struct {
	ExpenseRepo repositories.ExpenseRepository
	GroupRepo   repositories.GroupRepository
	DebtRepo    repositories.DebtRepository
}

func NewDebtService(
	expenseRepo repositories.ExpenseRepository,
	groupRepo repositories.GroupRepository,
	debtRepo repositories.DebtRepository,
) *DebtService {
	return &DebtService{
		ExpenseRepo: expenseRepo,
		GroupRepo:   groupRepo,
		DebtRepo:    debtRepo,
	}
}

func (s *DebtService) RecalculateDebts(groupID int) error {
	expenses, err := s.ExpenseRepo.GetExpensesByGroupID(groupID)
	if err != nil {
		return err
	}

	members, err := s.GroupRepo.GetGroupMembers(groupID)
	if err != nil {
		return err
	}

	debtMatrix := make(map[string]map[string]float64)

	for _, expense := range expenses {
		var debtors []string
		for _, member := range members {
			if member.ID != expense.CreatedBy {
				debtors = append(debtors, member.ID)
			}
		}

		if len(debtors) == 0 {
			continue
		}

		share := expense.Amount / float64(len(debtors))
		for _, debtor := range debtors {
			if _, ok := debtMatrix[debtor]; !ok {
				debtMatrix[debtor] = make(map[string]float64)
			}
			debtMatrix[debtor][expense.CreatedBy] += share
		}
	}

	var debts []models.Debt
	for from, toMap := range debtMatrix {
		for to, amount := range toMap {
			if amount > 0 {
				debts = append(debts, models.Debt{
					GroupID:    groupID,
					FromUserID: from,
					ToUserID:   to,
					Amount:     amount,
				})
			}
		}
	}

	if err := s.DebtRepo.ClearDebtsByGroupID(groupID); err != nil {
		return err
	}

	return s.DebtRepo.BulkInsertDebts(debts)
}

func (s *DebtService) GetMyDebtsInGroup(groupID int, userID string) ([]models.DebtDto, error) {
	return s.DebtRepo.GetUserDebtWithNames(groupID, userID)
}
