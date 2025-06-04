package controllers

import (
	"net/http"
	"strconv"

	"cospend/constant"
	"cospend/middleware"
	"cospend/models"
	"cospend/services"

	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
	ExpenseService services.ExpenseService
}

func NewExpenseController(service services.ExpenseService) *ExpenseController {
	return &ExpenseController{ExpenseService: service}
}

// CreateExpense godoc
// @Summary Создать расход в группе
// @Description Добавить новый расход в указанную группу
// @Tags expenses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Group ID"
// @Param expense body models.ExpenseCreateRequest true "Expense"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/groups/{id}/expenses [post]
func (ctrl *ExpenseController) CreateExpense(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code: http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: "Invalid group ID",
		})
		return
	}

	var req models.ExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code: http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
		})
		return
	}

	userClaims := middleware.GetUserClaims(c)

	expense := &models.Expense{
		GroupID:     groupID,
		CreatedBy:   userClaims.ID,
		Description: req.Description,
		Amount:      req.Amount,
		Date:        req.Date,
	}

	if err := ctrl.ExpenseService.CreateExpense(expense); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Code:         http.StatusCreated,
		ResponseCode: constant.SUCCESS,
		ResponseData: expense,
	})
}

// GetExpensesByGroup godoc
// @Summary Получить список расходов группы
// @Description Возвращает все расходы указанной группы
// @Tags expenses
// @Produce json
// @Security BearerAuth
// @Param id path int true "Group ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/groups/{id}/expenses [get]
func (ctrl *ExpenseController) GetExpensesByGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: "Invalid group ID",
		})
		return
	}

	expenses, err := ctrl.ExpenseService.GetExpensesByGroupID(groupID)
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
		ResponseData: expenses,
	})
}

// GetExpenseByID godoc
// @Summary Получить расход по ID
// @Description Возвращает информацию о конкретном расходе
// @Tags expenses
// @Produce json
// @Security BearerAuth
// @Param id path int true "Expense ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /api/expenses/{id} [get]
func (ctrl *ExpenseController) GetExpenseByID(c *gin.Context) {
	expenseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: "Invalid expense ID",
		})
		return
	}

	expense, err := ctrl.ExpenseService.GetExpenseByID(expenseID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:         http.StatusNotFound,
			ResponseCode: constant.FAILED_NOT_FOUND,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseData: expense,
	})
}
