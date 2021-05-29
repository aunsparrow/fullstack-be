package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=Asia/Bangkok", dbHost, dbPort, username, dbName, password) //Build connection string

	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbUri,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}

	Db = conn
	fmt.Println("Database connection successfully")
}

func OpenDB() *gorm.DB {
	return Db
}
