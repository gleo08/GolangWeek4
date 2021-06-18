package database

import (
	"fmt"

	"github.com/gleo08/GolangWeek4/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DbURL = "root:12345678@tcp(127.0.0.1:3306)/golangweek4?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() {
	database, err := gorm.Open(mysql.Open(DbURL), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	} else {
		fmt.Println("Successfully")
	}

	DB = database

	database.AutoMigrate(&models.Product{}, &models.User{}, &models.Cart{}, &models.Category{}, &models.Price{}, &models.Image{})
}
