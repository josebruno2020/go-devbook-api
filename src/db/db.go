package db

import (
	"log"

	"github.com/josebruno2020/go-devbook-api/src/config"
	"github.com/josebruno2020/go-devbook-api/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.ConnectionDB), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
