package routes

import (
	"cospend/controllers"
	"cospend/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	authController controllers.AuthController,
	userController controllers.UserController,
	groupController controllers.GroupController,
	expenseController controllers.ExpenseController,
	debtController controllers.DebtController,
	settlementController controllers.SettlementController,
) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	baseRouter := router.Group("/v1")

	/* ======= AUTH ======= */
	auth := baseRouter.Group("/auth")
	auth.POST("/login", authController.Login)

	// /* ======= USERS ======= */
	users := baseRouter.Group("/users")
	users.POST("", userController.CreateUser)

	/* ======= GROUPS ======= */
	groups := baseRouter.Group("/groups")
	groups.Use(middleware.AuthMiddleware())                     // Защита всех маршрутов JWT
	groups.POST("", groupController.CreateGroup)                // POST /v1/groups
	groups.GET("", groupController.GetUserGroups)               // GET /v1/groups
	groups.GET("/:id", groupController.GetGroupByID)            // GET /v1/groups/:id
	groups.POST("/:id/join", groupController.JoinGroup)         // POST /v1/groups/:id/join
	groups.GET("/:id/members", groupController.GetGroupMembers) // GET /v1/groups/:id/members

	/* ======= EXPENSES ======= */
	expenses := baseRouter.Group("/groups/:id/expenses")
	expenses.Use(middleware.AuthMiddleware())
	expenses.POST("", expenseController.CreateExpense)      // POST /v1/groups/:id/expenses
	expenses.GET("", expenseController.GetExpensesByGroup)  // GET /v1/groups/:id/expenses

	expense := baseRouter.Group("/expenses/:id")
	expense.Use(middleware.AuthMiddleware())
	expense.GET("", expenseController.GetExpenseByID)       // GET /v1/expenses/:id

	/* ======= DEBTS ======= */
	debt := baseRouter.Group("/groups/:id/debts")
	debt.Use(middleware.AuthMiddleware())
	debt.POST("distribute", debtController.RecalculateDebts)   // POST /v1/groups/:id/debts/distribute
	debt.GET("me", debtController.GetMyDebts) 			 	   // GET /v1/groups/:id/debts/me"

	/* ======= SETTLEMENTS ======= */
	settlement := baseRouter.Group("/groups/:id/settle")
	settlement.Use(middleware.AuthMiddleware())
	settlement.POST("", settlementController.SettleDebt) 	   // POST /v1/groups/:id/settle
	return router
}
