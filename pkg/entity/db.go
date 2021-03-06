package entity

import (
	"backend/config"
	"backend/internal/app/company/repository"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

var err error

var (
	Company = repository.NewCompanyRepository()
)

func Connect() (*gorm.DB, error) {
	//Connect to db using GORM
	host := config.Config.GetString("DB.ADDRESS")
	port := config.Config.GetString("DB.PORT")
	user := config.Config.GetString("DB.USERNAME")
	password := config.Config.GetString("DB.PASSWORD")
	dbName := config.Config.GetString("DB.DATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,
		},
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return DB, nil
}
