package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/Nuttachai-K/cafe-finder-api/internal/database"
)

func main() {
	// Load enviorment data
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect with database
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ctx := context.Background()

	// var id int
	// var name string

	// err = db.QueryRow(
	// 	ctx,
	// 	`SELECT id, name FROM cafes WHERE id = $1`,
	// 	1,
	// ).Scan(&id, &name)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(id, name)
}
