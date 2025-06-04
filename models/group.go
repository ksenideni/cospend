package models

import "time"

type Group struct {
	ID        int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" gorm:"type:text;not null"`
	CreatedBy string     `json:"createdBy,omitempty" gorm:"index"`
	CreatedAt time.Time  `json:"createdAt" gorm:"default:now()"`
}

func (Group) TableName() string {
	return "groups"
}

// DTOs

type GroupCreateRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type GroupResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}

type ListGroup struct {
	Page      int     `json:"page"`
	Limit     int     `json:"limit"`
	Total     int     `json:"total"`
	TotalPage int     `json:"totalPage"`
	Groups    []Group `json:"groups"`
}
