package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/api_go_gin"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	errMigrate := database.AutoMigrate(&Post{})
	if errMigrate != nil {
		return
	}
	DB = database
}
