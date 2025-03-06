package main

import (
	"log"

	"github.com/joao-vitor-felix/social-app-api/internal/env"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &application{
		config: config{
			address: env.GetString("PORT", ":8080"),
		},
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
