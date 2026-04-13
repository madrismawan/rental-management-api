package main

import (
	"log"

	"rental-management-api/config"
	"rental-management-api/internal/database"
	"rental-management-api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	db, err := database.New(cfg.Database.URL, cfg.Environment)
	if err != nil {
		log.Fatalf("database init failed: %v", err)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}

	if err := database.SeedAdminUser(db); err != nil {
		log.Fatalf("database seed failed: %v", err)
	}

	r := server.NewRouter(cfg, db)

	if err := r.Run(cfg.HTTP.Address()); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}
