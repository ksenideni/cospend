package controllers

import (
	"net/http"

	"cospend/constant"
	"cospend/middleware"
	"cospend/models"
	logger "cospend/pkg/logging"
	"cospend/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserRegister true "User"
// @Response 201 {object} models.Response
// @Response 500 {object} models.Response
// @Response 400 {object} models.Response
// @Response 302 {object} models.Response
// @Router /users [post]
func (controller *UserController) CreateUser(c *gin.Context) {
	var userRegister models.UserRegister
	if err := c.ShouldBindJSON(&userRegister); err != nil {
		middleware.Response(c, userRegister, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	user := models.User{
		Email:       userRegister.Email,
		Name:        userRegister.Name,
		PhoneNumber: userRegister.PhoneNumber,
		Password:    userRegister.Password,
		CreatedBy:   userRegister.Name,
	}

	response, err := controller.UserService.CreateUser(&user)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, userRegister, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		})
		return
	}

	middleware.Response(c, userRegister, *response)
}
