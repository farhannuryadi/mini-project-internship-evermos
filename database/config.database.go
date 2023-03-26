package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	godotenv.Load()
	dbhost	:= os.Getenv("MYSQL_HOST")
	dbuser	:= os.Getenv("MYSQL_USER")
	// dbpass	:= os.Getenv("MYSQL_PASSWORD")
	dbname	:= os.Getenv("MYSQL_DBNAME")
	
	dsn := fmt.Sprintf("%s:@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbhost, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to database")
	}
	fmt.Println("success connect to database")
	DB = db
}