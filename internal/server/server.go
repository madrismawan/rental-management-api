package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"rental-management-api/config"
	"rental-management-api/internal/handler"
	"rental-management-api/internal/repository"
	"rental-management-api/internal/service"
)

type Server struct {
	cfg    config.Config
	engine *gin.Engine
	logger zerolog.Logger
}

func NewRouter(cfg config.Config, db *gorm.DB) *gin.Engine {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	r.Use(gin.Recovery(), requestLogger(zerolog.Nop()))
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	registerRoutes(r, cfg, db)
	return r
}

func registerRoutes(engine *gin.Engine, cfg config.Config, db *gorm.DB) {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	api := engine.Group("/api/v1")

	userRepo := repository.NewUserRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	vehicleRepo := repository.NewVehicleRepository(db)
	rentalRepo := repository.NewRentalRepository(db)
	incidentRepo := repository.NewVehicleIncidentRepository(db)

	userSvc := service.NewUserService(db, userRepo)
	authSvc := service.NewAuthService(userSvc, cfg.Auth.AccessTokenSecret, cfg.Auth.RefreshTokenSecret, cfg.Auth.TokenTTL)
	customerSvc := service.NewCustomerService(db, userSvc, customerRepo)
	storageSvc := service.NewStorageService(cfg.Storage)
	vehicleSvc := service.NewVehicleService(db, vehicleRepo)
	rentalSvc := service.NewRentalService(db, rentalRepo, vehicleSvc)
	incidentSvc := service.NewVehicleIncidentService(db, incidentRepo)

	handler.NewAuthHandler(authSvc).RegisterRoutes(api)
	handler.NewUserHandler(userSvc).Register(api)
	handler.NewCustomerHandler(customerSvc, storageSvc).Register(api)
	handler.NewVehicleHandler(vehicleSvc).Register(api)
	handler.NewRentalHandler(rentalSvc).Register(api)
	handler.NewVehicleIncidentHandler(incidentSvc).Register(api)
}
