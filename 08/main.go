package main

import (
	"fmt"

	"gophercises/08/db"
	"gophercises/08/db/models"
	"gophercises/08/normalize"
)

func main() {
	DB := db.StartDB()
	defer DB.Close()
	db.Migrate(DB)

	phones := []models.Phone{}
	DB.Find(&phones)

	for _, phone := range phones {
		normalized := normalize.Normalize(phone.Number)
		DB.Model(&phone).Update("number", normalized)
	}

	phones = []models.Phone{}
	DB.Find(&phones)

	for _, phone := range phones {
		pSlice := []models.Phone{}
		DB.Where("number = ?", phone.Number).Find(&pSlice)

		if len(pSlice) > 1 {
			for i := 1; i < len(pSlice); i++ {
				DB.Delete(&pSlice[i])
			}
		}
	}

	phones = []models.Phone{}
	DB.Find(&phones)

	strSlice := []string{}
	for _, v := range phones {
		strSlice = append(strSlice, v.Number)
	}
	fmt.Println(strSlice)
}
