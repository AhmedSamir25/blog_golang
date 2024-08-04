package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbConn *gorm.DB // تأكد من أن الاسم هنا هو DbConn

func ConnDB() {
	dsn := "root:ahmed@sql123@tcp(localhost:3306)/blogs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("Database Connection Failed")
	}
	DbConn = db
}
