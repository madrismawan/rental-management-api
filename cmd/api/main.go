package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"rental-management-api/config"
	"rental-management-api/internal/database"
	"rental-management-api/internal/handler"
	"rental-management-api/internal/repository"
	"rental-management-api/internal/service"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("database init failed: %v", err)
	}

	r := gin.Default()

	api := r.Group("/api/v1")

	userRepo := repository.NewUserRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	vehicleRepo := repository.NewVehicleRepository(db)
	rentalRepo := repository.NewRentalRepository(db)
	incidentRepo := repository.NewVehicleIncidentRepository(db)

	userSvc := service.NewUserService(userRepo)
	customerSvc := service.NewCustomerService(customerRepo)
	vehicleSvc := service.NewVehicleService(vehicleRepo)
	rentalSvc := service.NewRentalService(rentalRepo)
	incidentSvc := service.NewVehicleIncidentService(incidentRepo)

	handler.NewUserHandler(userSvc).Register(api)
	handler.NewCustomerHandler(customerSvc).Register(api)
	handler.NewVehicleHandler(vehicleSvc).Register(api)
	handler.NewRentalHandler(rentalSvc).Register(api)
	handler.NewVehicleIncidentHandler(incidentSvc).Register(api)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}
