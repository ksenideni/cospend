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

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param auth body models.AuthLogin true "Auth"
// @Response 200 {object} models.Response
// @Response 500 {object} models.Response
// @Response 400 {object} models.Response
// @Router /auth/login [post]
func (controller *AuthController) Login(c *gin.Context) {
	var authLogin models.AuthLogin
	if err := c.ShouldBindJSON(&authLogin); err != nil {
		middleware.Response(c, authLogin, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	auth := models.AuthLogin{
		Email:    authLogin.Email,
		Password: authLogin.Password,
	}

	response, err := controller.AuthService.Login(&auth)
	if err != nil {
		logger.Err(err.Error())
		c.JSON(500, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	middleware.Response(c, authLogin, *response)
}
