package main

import (
	"Assigment_8/db"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db.SetupDB()
	err = db.DB.AutoMigrate(&db.WindWaterStatus{})
	if err != nil {
		log.Fatal("Error Migrate")
	} else {
		log.Println("Success Migrate")
	}
}
