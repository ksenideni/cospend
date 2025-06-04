package services

import (
	"cospend/models"
	"cospend/repositories"
	"errors"
)

type SettlementService struct {
	SettlementRepo repositories.SettlementRepository
}

func NewSettlementService(repo repositories.SettlementRepository) *SettlementService {
	return &SettlementService{SettlementRepo: repo}
}

func (s *SettlementService) SettleDebt(groupID int, fromUserID string, req models.SettleRequest) error {
	debt, err := s.SettlementRepo.GetDebt(groupID, fromUserID, req.ToUserID)
	if err != nil {
		return errors.New("долг не найден или уже погашен")
	}

	if req.Amount > debt.Amount {
		return errors.New("сумма превышает задолженность")
	}

	settlement := &models.Settlement{
		GroupID:    groupID,
		FromUserID: fromUserID,
		ToUserID:   req.ToUserID,
		Amount:     req.Amount,
	}

	if err := s.SettlementRepo.AddSettlement(settlement); err != nil {
		return err
	}

	return s.SettlementRepo.UpdateOrDeleteDebt(debt, req.Amount)
}
