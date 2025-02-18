package database

import (
	"go-auth/modles"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {

	// load the .env file
	// dsn := "root:newpassword@tcp(127.0.0.1:3306)/userDB?charset=utf8mb4&parseTime=True&loc=Local"

	// Load environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbParseTime := os.Getenv("DB_PARSE_TIME")
	dbLoc := os.Getenv("DB_LOC")

	// Build DSN (Data Source Name) dynamically
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName +
		"?charset=" + dbCharset + "&parseTime=" + dbParseTime + "&loc=" + dbLoc

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
