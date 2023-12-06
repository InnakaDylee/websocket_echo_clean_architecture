package config

import (
	"ws/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {


	dbUser := 
	dbPass := 
	dbHost := 
	dbPort := 
	dbName := 

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		panic("Failed to Connect Database")
	}

	Migrate()

	fmt.Println("Connected to Database")

	return DB
}

func Migrate() {
	DB.AutoMigrate(&model.User{}, &model.Rooms{}, &model.Message{})
}
