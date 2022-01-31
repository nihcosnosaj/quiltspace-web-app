package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	// load database config string as environment variable
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dotenv := os.Getenv("psqlconfig")

	// var err error
	DB, err = sql.Open("postgres", dotenv)
	if err != nil {
		log.Fatalln(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("You connected to your database.")
}
