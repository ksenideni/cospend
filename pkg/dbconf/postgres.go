package dbconf

import (
	"context"
	"log"

	"cospend/pkg/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbGorm *gorm.DB

func InitGorm(ctx context.Context) (*gorm.DB, error) {
	var err error
	// connectionString := util.Getenv("GORM_CONNECTION")
	connectionString := util.Getenv("LOCAL_GORM_CONNECTION")
	DbGorm, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	return DbGorm, nil
}
