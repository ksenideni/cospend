package models

import (
	"time"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID          string     `json:"id" gorm:"primary_key;not null;type:varchar(100);index"`
	Email       string     `json:"email" gorm:"type:varchar(100);not null"`
	PhoneNumber string     `json:"phoneNumber" gorm:"type:varchar(20);not null"`
	Name        string     `json:"name" gorm:"type:varchar(250);not null"`
	Password    string     `json:"password,omitempty" gorm:"type:varchar(150)"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"not null;default:now()"`
	CreatedBy   string     `json:"createdBy" gorm:"type:varchar(150);not null"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	UpdatedBy   *string    `json:"updatedBy,omitempty" gorm:"type:varchar(150)"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
	DeletedBy   *string    `json:"deletedBy,omitempty" gorm:"type:varchar(150)"`
}

func (User) TableName() string {
	return "users"
}

// DTOs

type ListUser struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Total    int    `json:"total"`
	Users    []User `json:"users"`
	TotalPage int   `json:"totalPage"`
}

type UserRegister struct {
	Email       string `json:"email" binding:"required,email"`
	Name        string `json:"name" binding:"required,min=3"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type UserDto struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required,min=3"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}


type UserClaims struct {
	jwt.StandardClaims
	ID          string `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}