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

type GroupController struct {
	GroupService services.GroupService
}

func NewGroupController(service services.GroupService) *GroupController {
	return &GroupController{GroupService: service}
}

// CreateGroup godoc
// @Summary Создание группы
// @Description Создает новую группу и добавляет текущего пользователя в участники
// @Tags groups
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param group body models.GroupCreateRequest true "Group"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/groups [post]
func (ctrl *GroupController) CreateGroup(c *gin.Context) {
	var req models.GroupCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, ResponseCode: constant.FAILED_REQUIRED, ResponseDesc: err.Error()})
		return
	}

	userClaims := middleware.GetUserClaims(c)
	group := &models.Group{
		Name:      req.Name,
		CreatedBy: userClaims.ID,
	}

	err := ctrl.GroupService.CreateGroup(group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: http.StatusInternalServerError, ResponseCode: constant.FAILED_INTERNAL})
		return
	}

	// Автоматически добавляем создателя в участники
	_ = ctrl.GroupService.JoinGroup(group.ID, userClaims.ID)

	c.JSON(http.StatusCreated, models.Response{Code: http.StatusCreated, ResponseCode: constant.SUCCESS, ResponseData: group})
}

// GetUserGroups godoc
// @Summary Получить список групп пользователя
// @Description Возвращает все группы, в которых участвует текущий пользователь
// @Tags groups
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/groups [get]
func (ctrl *GroupController) GetUserGroups(c *gin.Context) {
	userClaims := middleware.GetUserClaims(c)

	groups, err := ctrl.GroupService.GetUserGroups(userClaims.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: http.StatusInternalServerError, ResponseCode: constant.FAILED_INTERNAL})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, ResponseCode: constant.SUCCESS, ResponseData: groups})
}

// GetGroupByID godoc
// @Summary Получить информацию о группе
// @Description Возвращает информацию о конкретной группе по ID
// @Tags groups
// @Produce json
// @Security BearerAuth
// @Param id path int true "Group ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /api/groups/{id} [get]
func (ctrl *GroupController) GetGroupByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	group, err := ctrl.GroupService.GetGroup(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{Code: http.StatusNotFound, ResponseCode: constant.FAILED_NOT_FOUND})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, ResponseCode: constant.SUCCESS, ResponseData: group})
}

// JoinGroup godoc
// @Summary Присоединиться к группе
// @Description Добавляет текущего пользователя в список участников группы
// @Tags groups
// @Produce json
// @Security BearerAuth
// @Param id path int true "Group ID"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/groups/{id}/join [post]
func (ctrl *GroupController) JoinGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userClaims := middleware.GetUserClaims(c)

	err := ctrl.GroupService.JoinGroup(id, userClaims.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: http.StatusInternalServerError, ResponseCode: constant.FAILED_INTERNAL})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, ResponseCode: constant.SUCCESS, ResponseDesc: "Joined group"})
}

// GetGroupMembers godoc
// @Summary Получить список участников группы
// @Description Возвращает всех пользователей, присоединившихся к указанной группе
// @Tags groups
// @Produce json
// @Security BearerAuth
// @Param id path int true "Group ID"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/groups/{id}/members [get]
func (ctrl *GroupController) GetGroupMembers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := ctrl.GroupService.GetGroupMembers(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: http.StatusInternalServerError, ResponseCode: constant.FAILED_INTERNAL})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, ResponseCode: constant.SUCCESS, ResponseData: users})
}
