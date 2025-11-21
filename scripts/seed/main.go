package main

import (
	"log"

	"flynt/internal/database"
	"flynt/internal/utils"
)

func main() {
	err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	cfg := utils.GetConfig()

	db, err := database.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	err = db.SeedData();
	if err != nil {
		log.Printf("Error inserting dummy data: %v", err)
	}
}
