package config

import (
	"app/model"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	godotenv.Load(".env")
	//err := godotenv.Load(filepath.Join(".", ".env"))
	//if err != nil {
	//	fmt.Println("Error loading .env file")
	//	os.Exit(1)
	//}

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		panic("Failed to Connect Database")
	}

	InitMigrate()

	fmt.Println("Connected to Database")
}

func InitMigrate() {
	DB.AutoMigrate(
		&model.User{},
		&model.News{},
		&model.Contest{},
		&model.Contestant{},
	)
}
