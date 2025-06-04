package services

import (
	"cospend/models"
	"cospend/repositories"
)

type GroupService struct {
	GroupRepo repositories.GroupRepository
}

func NewGroupService(groupRepo repositories.GroupRepository) *GroupService {
	return &GroupService{GroupRepo: groupRepo}
}

func (s *GroupService) CreateGroup(group *models.Group) error {
	return s.GroupRepo.CreateGroup(group)
}

func (s *GroupService) GetUserGroups(userID string) ([]models.Group, error) {
	return s.GroupRepo.GetGroupsByUserID(userID)
}

func (s *GroupService) GetGroup(groupID int) (models.Group, error) {
	return s.GroupRepo.GetGroupByID(groupID)
}

func (s *GroupService) JoinGroup(groupID int, userID string) error {
	return s.GroupRepo.JoinGroup(groupID, userID)
}

func (s *GroupService) GetGroupMembers(groupID int) ([]models.UserDto, error) {
	users, err := s.GroupRepo.GetGroupMembers(groupID)
	if err != nil {
		return nil, err
	}

	var userDtos []models.UserDto
	for _, u := range users {
		userDtos = append(userDtos, models.UserDto{
			ID:          u.ID,
			Name:        u.Name,
			Email:       u.Email,
			PhoneNumber: u.PhoneNumber,
		})
	}

	return userDtos, nil
}
