package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var host = os.Getenv("DB_HOST")
var username = os.Getenv("DB_USERNAME")
var password = os.Getenv("DB_PASSWORD")
var port = os.Getenv("DB_PORT")
var dbName = os.Getenv("DB_NAME")

var dsn = "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable"
var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Info),
})

func GetDB() (*gorm.DB, error) {
	if err != nil {
		return nil, err
	}
	return db, nil
}
