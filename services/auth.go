package services

import (
	"net/http"

	"cospend/constant"
	"cospend/middleware"
	"cospend/pkg/util"

	"cospend/models"
	"cospend/repositories"
)

type AuthService struct {
	UserRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Login(auth *models.AuthLogin) (res *models.Response, err error) {
	authUser, err := service.UserRepository.GetUserByEmailOrPhoneNumber(auth.Email, auth.Email)
	if err != nil {
		return &models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: "Email or Phone Number not valid",
		}, nil
	}

	err = util.CheckPassword(auth.Password, authUser.Password)
	if err != nil {
		return &models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_AUTHORIZED,
			ResponseDesc: "Password not valid",
		}, nil
	}

	userClaims := models.UserClaims{
		ID:          authUser.ID,
		Name:        authUser.Name,
		Email:       authUser.Email,
		PhoneNumber: authUser.PhoneNumber,
	}

	token, err := middleware.GenerateToken(userClaims)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: "User logged in successfully",
		ResponseData: models.AuthToken{
			Token: token,
		},
	}, nil
}
