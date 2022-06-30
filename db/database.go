package db

import (
	"fmt"
	"log"
	"os"

	"github.com/TulioGuaraldoB/school-report/db/migration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb(
	DbUser string,
	DbPassword string,
	DbHost string,
	DbPort string,
	DbName string,
) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		DbName,
	)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errMessage := fmt.Sprintf("Failed to connect do MySql %s", err.Error())
		log.Fatal(errMessage)
	}

	fmt.Println("Connected to MySql!")

	db = database

	migration.Run(db)
}

func OpenConnection() *gorm.DB {
	StartDb(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	return db
}
