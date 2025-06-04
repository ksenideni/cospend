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

type DebtController struct {
	DebtService services.DebtService
}

func NewDebtController(service services.DebtService) *DebtController {
	return &DebtController{DebtService: service}
}

// POST /api/groups/:id/debts/distribute
func (ctrl *DebtController) RecalculateDebts(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
		})
		return
	}

	err = ctrl.DebtService.RecalculateDebts(groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: "Долги перерасчитаны",
	})
}

// GET /api/groups/:id/debts/me
func (ctrl *DebtController) GetMyDebts(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
		})
		return
	}

	user := middleware.GetUserClaims(c)

	debts, err := ctrl.DebtService.GetMyDebtsInGroup(groupID, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseData: debts,
	})
}
