package repositories

import (
	"cospend/models"
	"gorm.io/gorm"
)

type DebtRepository struct {
	DB *gorm.DB
}

func NewDebtRepository(db *gorm.DB) *DebtRepository {
	return &DebtRepository{DB: db}
}

func (r *DebtRepository) ClearDebtsByGroupID(groupID int) error {
	return r.DB.Where("group_id = ?", groupID).Delete(&models.Debt{}).Error
}

func (r *DebtRepository) BulkInsertDebts(debts []models.Debt) error {
	for _, d := range debts {
		err := r.DB.Exec(`
			INSERT INTO debts (group_id, from_user_id, to_user_id, amount)
			VALUES (?, ?, ?, ?)
			ON CONFLICT (group_id, from_user_id, to_user_id)
			DO UPDATE SET amount = EXCLUDED.amount
		`, d.GroupID, d.FromUserID, d.ToUserID, d.Amount).Error

		if err != nil {
			return err
		}
	}
	return nil
}

func (r *DebtRepository) GetUserDebtWithNames(groupID int, userID string) ([]models.DebtDto, error) {
	var dtos []models.DebtDto
	err := r.DB.Raw(`
		SELECT d.to_user_id, u.name AS to_user_name, d.amount
		FROM debts d
		JOIN users u ON u.id = d.to_user_id
		WHERE d.group_id = ? AND d.from_user_id = ?
	`, groupID, userID).Scan(&dtos).Error
	return dtos, err
}
