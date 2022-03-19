package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

// init before exec main
func init() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot load file .env: ", err)
		panic(err)
	}
	log.Println("ENV:", os.Getenv("MODE"))
}