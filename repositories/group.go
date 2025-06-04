package repositories

import (
	"cospend/models"

	"gorm.io/gorm"
)

type GroupRepository struct {
	DB *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{DB: db}
}

func (r *GroupRepository) CreateGroup(group *models.Group) error {
	return r.DB.Create(group).Error
}

func (r *GroupRepository) GetGroupsByUserID(userID string) ([]models.Group, error) {
	var groups []models.Group
	err := r.DB.Raw(`
		SELECT g.*
		FROM groups g
		JOIN group_members gm ON g.id = gm.group_id
		WHERE gm.user_id = ?`, userID).Scan(&groups).Error
	return groups, err
}

func (r *GroupRepository) GetGroupByID(groupID int) (models.Group, error) {
	var group models.Group
	err := r.DB.First(&group, groupID).Error
	return group, err
}

func (r *GroupRepository) JoinGroup(groupID int, userID string) error {
	return r.DB.Exec(`
		INSERT INTO group_members (group_id, user_id)
		VALUES (?, ?)
		ON CONFLICT (group_id, user_id) DO NOTHING
	`, groupID, userID).Error
}

func (r *GroupRepository) GetGroupMembers(groupID int) ([]models.User, error) {
	var users []models.User
	err := r.DB.Raw(`
		SELECT u.*
		FROM users u
		JOIN group_members gm ON u.id = gm.user_id
		WHERE gm.group_id = ?`, groupID).Scan(&users).Error
	return users, err
}
