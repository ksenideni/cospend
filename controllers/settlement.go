package controllers

import (
	"cospend/constant"
	"cospend/middleware"
	"cospend/models"
	"cospend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SettlementController struct {
	Service *services.SettlementService
}

func NewSettlementController(service *services.SettlementService) *SettlementController {
	return &SettlementController{Service: service}
}

// POST /api/groups/:id/settle
func (ctrl *SettlementController) SettleDebt(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, ResponseCode: constant.FAILED_REQUIRED})
		return
	}

	var req models.SettleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, ResponseCode: constant.FAILED_REQUIRED, ResponseDesc: err.Error()})
		return
	}

	user := middleware.GetUserClaims(c)

	err = ctrl.Service.SettleDebt(groupID, user.ID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: "Долг погашен",
	})
}
