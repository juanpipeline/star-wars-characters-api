package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/juanpipeline/star-wars-characters-api/models"
)

func GetConnection() *gorm.DB {

	Host := os.Getenv("DB_HOST")
	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASSWORD")
	Dbname := os.Getenv("DB_NAME")

	connection_string := "host=" + Host + " user=" + User + " password=" + Password + " dbname=" + Dbname + " port=5432 sslmode=disable"
	db, error := gorm.Open(postgres.Open(connection_string), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	log.Println("Automigration working....")

	db.AutoMigrate(&models.Character{})
}
