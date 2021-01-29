package orm

import (
	"api/src/database"
	"api/src/models"
	"log"
)

func AutoMigration(){
	db, err := database.GetBD()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Users{},&models.Posts{})
}
