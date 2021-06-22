package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"root/models"
)

var DB *gorm.DB

func Connect() {
	//db
	//mysqlUser:="root"
	//mysqlPass:="rootroot"
	Connection, err := gorm.Open(mysql.Open("root:49038047@/snapfood_go_auth"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = Connection
	Connection.AutoMigrate(&models.Manager{}, &models.User{}, &models.Food{}, &models.Restaurant{})
}
