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

}
