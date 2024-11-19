package config

import (
	"fmt"
	"log"
	"test-ottodigital-be/domain/model"
	"test-ottodigital-be/utils"

	"github.com/op/go-logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = logging.MustGetLogger("main")

func ConnectDatabase() *gorm.DB {
	var (
		HOST    = utils.GetEnv("DB_HOST", "localhost")
		USER    = utils.GetEnv("DB_USER", "postgres")
		PORT    = utils.GetEnv("DB_PORT", "5432")
		DB_NAME = utils.GetEnv("DB_NAME", "test")
	)

	uri := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", HOST, USER, DB_NAME, PORT)
	dsn := uri
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	db.AutoMigrate(&model.Brand{})
	db.AutoMigrate(&model.Voucher{})
	db.AutoMigrate(&model.VoucherRedemption{})

	return db
}
