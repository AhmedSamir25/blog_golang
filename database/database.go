package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbConn *gorm.DB // تأكد من أن الاسم هنا هو DbConn

func ConnDB() {
	// تحميل القيم من ملف .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// قراءة القيم من البيئة
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// بناء سلسلة الاتصال بقاعدة البيانات
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// فتح الاتصال بقاعدة البيانات
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("Database Connection Failed")
	}
	DbConn = db
}
