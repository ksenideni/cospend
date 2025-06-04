package models

import "time"

type GroupMember struct {
	ID       int       `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID  int       `json:"groupId" gorm:"not null;index"`
	UserID   string    `json:"userId" gorm:"not null;index"`
	JoinedAt time.Time `json:"joinedAt" gorm:"default:now()"`
}

func (GroupMember) TableName() string {
	return "group_members"
}

type GroupMemberAdd struct {
	GroupID int    `json:"groupId" binding:"required"`
	UserID  string `json:"userId" binding:"required"`
}

type ListGroupMember struct {
	Page         int           `json:"page"`
	Limit        int           `json:"limit"`
	Total        int           `json:"total"`
	TotalPage    int           `json:"totalPage"`
	GroupMembers []GroupMember `json:"groupMembers"`
}
