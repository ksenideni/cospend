package repositories

import (
	"cospend/models"
	"gorm.io/gorm"
)

type SettlementRepository struct {
	DB *gorm.DB
}

func NewSettlementRepository(db *gorm.DB) *SettlementRepository {
	return &SettlementRepository{DB: db}
}

func (r *SettlementRepository) AddSettlement(settlement *models.Settlement) error {
	return r.DB.Create(settlement).Error
}

func (r *SettlementRepository) GetDebt(groupID int, fromUserID, toUserID string) (*models.Debt, error) {
	var debt models.Debt
	err := r.DB.
		Where("group_id = ? AND from_user_id = ? AND to_user_id = ?", groupID, fromUserID, toUserID).
		First(&debt).Error
	if err != nil {
		return nil, err
	}
	return &debt, nil
}

func (r *SettlementRepository) UpdateOrDeleteDebt(debt *models.Debt, paidAmount float64) error {
	if debt.Amount <= paidAmount {
		return r.DB.Delete(debt).Error
	}

	debt.Amount -= paidAmount
	return r.DB.Save(debt).Error
}
