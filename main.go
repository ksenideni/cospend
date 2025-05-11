package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"cospend/db"
	"cospend/util"
)

func main() {
	// загрузить переменные среды
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	databaseURL := util.Getenv("GORM_CONNECTION")

	// Подключение к БД
	// dbConn, err := sql.Open("postgres", databaseURL)
	// if err != nil {
	//     log.Fatal("DB connection error:", err)
	// }

	// === МИГРАЦИИ ===
	if err := db.RunMigrations(databaseURL); err != nil {
		log.Fatal("Migration failed:", err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	r.Run(":8080")
}
