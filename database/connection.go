package database

import (
	"fmt"
	"go-auth/modles"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {

	// load the .env file

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get database variables form .env file
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbParseTime := os.Getenv("DB_PARSE_TIME")
	dbLoc := os.Getenv("DB_LOC")

	// dsn := "root:newpassword@tcp(127.0.0.1:3306)/userDB?charset=utf8mb4&parseTime=True&loc=Local"
	// Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset, dbParseTime, dbLoc,
	)
	conDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	DB = conDB

	if err := conDB.AutoMigrate(&modles.User{}); err != nil {
		log.Fatal("Auto Migrate Failed User", err)
	}

	if err := conDB.AutoMigrate(&modles.Product{}); err != nil {
		log.Fatal("Auto Migrate Failed Product", err)
	}

}
