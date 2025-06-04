package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"cospend/controllers"
	"cospend/pkg/dbconf"
	"cospend/pkg/migrator"
	"cospend/pkg/util"
	"cospend/repositories"
	"cospend/routes"
	"cospend/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// загрузить переменные среды
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// databaseURL := util.Getenv("GORM_CONNECTION")
	databaseURL := util.Getenv("LOCAL_GORM_CONNECTION")

	// Подключение к БД
	// dbConn, err := sql.Open("postgres", databaseURL)
	// if err != nil {
	//     log.Fatal("DB connection error:", err)
	// }

	// === МИГРАЦИИ ===
	if err := migrator.RunMigrations(databaseURL); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// инициализация gorm
	ctx := context.Background()

	db, err := dbconf.InitGorm(ctx)
	if err != nil {
		panic(err)
	}

	// Repositories
	userRepository := repositories.NewUserRepository(db)
	groupRepository := repositories.NewGroupRepository(db)
	expenseRepository := repositories.NewExpenseRepository(db)
	debtRepository := repositories.NewDebtRepository(db)
	settlementRepository := repositories.NewSettlementRepository(db)

	// Services
	authService := services.NewAuthService(*userRepository)
	userService := services.NewUserService(*userRepository)
	groupService := services.NewGroupService(*groupRepository)
	expenseService := services.NewExpenseService(*expenseRepository)
	debtService := services.NewDebtService(*expenseRepository, *groupRepository, *debtRepository)
	settlementService := services.NewSettlementService(*settlementRepository)

	// Controllers
	authController := controllers.NewAuthController(*authService)
	userController := controllers.NewUserController(*userService)
	groupController := controllers.NewGroupController(*groupService)
	expenseController := controllers.NewExpenseController(*expenseService)
	debtController := controllers.NewDebtController(*debtService)
	settlementController := controllers.NewSettlementController(*&settlementService)

	router := routes.NewRouter(
		*authController, 
		*userController, 
		*groupController,
		*expenseController,
		*debtController,
		*settlementController,
	)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := fmt.Sprintf(":%s", util.Getenv("HTTP_PORT"))
	router.Run(port)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	r.Run(":8080")
}
