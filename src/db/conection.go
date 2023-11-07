package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "host=localhost user=postgres password=1008 dbname=prueba port=5432"
var DB *gorm.DB

func DbConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if error != nil {
		log.Println("Failed to connect to database!")
	} else {
		log.Println("Se conecto a la base de datos ")
	}

}
