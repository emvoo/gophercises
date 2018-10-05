package db

import (
	"log"

	"gophercises/08/db/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func StartDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "my.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Migrate(db *gorm.DB) {
	db = db.AutoMigrate(&models.Phone{})

	numbers := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	for _, p := range numbers {
		db.Create(&models.Phone{Number: p})
	}
}
