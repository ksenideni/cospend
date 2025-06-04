package repositories

import (
	"cospend/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) AddUser(user *models.User) error {
	return repository.DB.Create(user).Error
}

func (repository *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := repository.DB.Where(&models.User{Email: email}).Where("deleted_at isnull").First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repository *UserRepository) GetUserByPhoneNumber(phoneNumber string) (models.User, error) {
	var user models.User
	if err := repository.DB.Where(&models.User{PhoneNumber: phoneNumber}).Where("deleted_at isnull").First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repository *UserRepository) GetUserByEmailOrPhoneNumber(email, phoneNumber string) (models.User, error) {
	var user models.User
	if err := repository.DB.Where(&models.User{Email: email}).Or(&models.User{PhoneNumber: phoneNumber}).Where("deleted_at isnull").First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
