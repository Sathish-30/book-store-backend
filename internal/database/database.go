package database

import (
	"fmt"
	"log"

	"github.com/sathish-30/book-management/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// "host=localhost user=gorm password=gorm dbname=gorm port=9920

var dsn string = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", config.Configuration.HOST, config.Configuration.DBUserName, config.Configuration.DBUserPassword, config.Configuration.DBName, config.Configuration.DBPort)
var Db *gorm.DB

func ConnectToDatabase() {
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error -> ", err)
	}
	log.Println("Database connected successfully")
}
