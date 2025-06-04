package services

import (
	"net/http"

	"cospend/constant"
	logger "cospend/pkg/logging"
	"cospend/pkg/util"

	"cospend/models"
	"cospend/repositories"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (service *UserService) CreateUser(user *models.User) (res *models.Response, err error) {
	_, err = service.UserRepository.GetUserByEmail(user.Email)
	if err == nil {
		return &models.Response{
			Code:         http.StatusFound,
			ResponseCode: constant.FAILED_EXIST,
			ResponseDesc: "Email already exist",
		}, nil
	}

	_, err = service.UserRepository.GetUserByPhoneNumber(user.PhoneNumber)
	if err == nil {
		return &models.Response{
			Code:         http.StatusFound,
			ResponseCode: constant.FAILED_EXIST,
			ResponseDesc: "Phone Number already exist",
		}, nil
	}

	password, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.ID = uuid.New().String()
	logger.Infof("new user id"+user.ID)
	user.Password = password
	err = service.UserRepository.AddUser(user)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:         http.StatusCreated,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: "User created successfully",
	}, nil
}
