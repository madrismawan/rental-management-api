package database

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"rental-management-api/internal/constant"
	"rental-management-api/internal/entity"
)

const (
	seedAdminName               = "Administrator"
	seedAdminEmail              = "admin@gmail.com"
	defaultSeedAdminPassword    = "admin123"
	defaultSeedCustomerPassword = "customer123"
	seedCustomerTarget          = 10
	seedVehicleTarget           = 100
	seedRentalTarget            = 10
	seedIncidentTarget          = 15
)

func SeedAll(db *gorm.DB) error {
	if err := SeedAdminUser(db); err != nil {
		return fmt.Errorf("seed admin user: %w", err)
	}
	if err := SeedCustomers(db); err != nil {
		return fmt.Errorf("seed customers: %w", err)
	}
	if err := SeedVehicles(db); err != nil {
		return fmt.Errorf("seed vehicles: %w", err)
	}
	if err := SeedRentals(db); err != nil {
		return fmt.Errorf("seed rentals: %w", err)
	}
	if err := SeedVehicleIncidents(db); err != nil {
		return fmt.Errorf("seed vehicle incidents: %w", err)
	}

	return nil
}

func SeedCustomers(db *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultSeedCustomerPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash customer password: %w", err)
	}

	for i := 1; i <= seedCustomerTarget; i++ {
		email := fmt.Sprintf("customer%03d@seed.local", i)
		name := fmt.Sprintf("Customer %03d", i)
		phone := fmt.Sprintf("08120000%04d", i)

		var user entity.User
		err := db.Where("email = ?", email).First(&user).Error
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("query customer user %s: %w", email, err)
			}

			user = entity.User{
				Name:     name,
				Email:    email,
				Role:     constant.UserRoleCustomer,
				Password: string(hashedPassword),
			}
			if createErr := db.Create(&user).Error; createErr != nil {
				return fmt.Errorf("create customer user %s: %w", email, createErr)
			}
		} else if user.Role != constant.UserRoleCustomer {
			if updateErr := db.Model(&user).Update("role", constant.UserRoleCustomer).Error; updateErr != nil {
				return fmt.Errorf("update customer role %s: %w", email, updateErr)
			}
		}

		var customer entity.Customer
		customerErr := db.Where("user_id = ?", user.ID).First(&customer).Error
		if customerErr != nil {
			if !errors.Is(customerErr, gorm.ErrRecordNotFound) {
				return fmt.Errorf("query customer profile %s: %w", email, customerErr)
			}

			newCustomer := entity.Customer{
				UserID:      user.ID,
				PhoneNumber: phone,
				Address:     "Seed Address",
				Status:      constant.CustomerStatusActive,
				AvatarURL:   "",
			}
			if createErr := db.Create(&newCustomer).Error; createErr != nil {
				return fmt.Errorf("create customer profile %s: %w", email, createErr)
			}
		} else if customer.Status == "" {
			if updateErr := db.Model(&customer).Update("status", constant.CustomerStatusActive).Error; updateErr != nil {
				return fmt.Errorf("set customer status %s: %w", email, updateErr)
			}
		}
	}

	return nil
}

func SeedAdminUser(db *gorm.DB) error {
	var user entity.User
	err := db.Where("email = ?", seedAdminEmail).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("query admin user: %w", err)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		adminPassword := os.Getenv("ADMIN_PASSWORD")
		if adminPassword == "" {
			adminPassword = defaultSeedAdminPassword
		}

		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		if hashErr != nil {
			return fmt.Errorf("hash admin password: %w", hashErr)
		}

		admin := entity.User{
			Name:     seedAdminName,
			Email:    seedAdminEmail,
			Role:     constant.UserRoleAdmin,
			Password: string(hashedPassword),
		}
		if createErr := db.Create(&admin).Error; createErr != nil {
			return fmt.Errorf("create admin user: %w", createErr)
		}
		return nil
	}

	if user.Role != constant.UserRoleAdmin {
		if updateErr := db.Model(&user).Update("role", constant.UserRoleAdmin).Error; updateErr != nil {
			return fmt.Errorf("update admin role: %w", updateErr)
		}
	}

	return nil
}

func SeedVehicles(db *gorm.DB) error {
	var count int64
	if err := db.Model(&entity.Vehicle{}).Count(&count).Error; err != nil {
		return fmt.Errorf("count vehicles: %w", err)
	}

	if count >= seedVehicleTarget {
		return nil
	}

	brands := []string{"Toyota", "Honda", "Suzuki", "Mitsubishi", "Daihatsu", "Nissan", "Mazda", "Hyundai", "Kia", "Wuling"}
	models := []string{"Avanza", "Xenia", "Brio", "Civic", "Ertiga", "Xpander", "Sigra", "Livina", "CX-5", "Almaz"}
	colors := []string{"Black", "White", "Silver", "Gray", "Red", "Blue"}
	statuses := []constant.VehicleStatus{
		constant.VehicleStatusAvailable,
		constant.VehicleStatusAvailable,
		constant.VehicleStatusAvailable,
		constant.VehicleStatusRented,
		constant.VehicleStatusMaintenance,
	}
	conditions := []constant.VehicleCondition{
		constant.VehicleConditionGood,
		constant.VehicleConditionGood,
		constant.VehicleConditionService,
		constant.VehicleConditionBroke,
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	start := int(count) + 1
	toCreate := seedVehicleTarget - int(count)
	vehicles := make([]entity.Vehicle, 0, toCreate)

	for i := 0; i < toCreate; i++ {
		n := start + i
		plate := fmt.Sprintf("B %04d GX", n)
		brand := brands[rng.Intn(len(brands))]
		model := models[rng.Intn(len(models))]
		color := colors[rng.Intn(len(colors))]
		year := 2016 + rng.Intn(10)
		cc := 1000 + rng.Intn(2001)
		mileage := 10000 + rng.Intn(140001)
		dailyRate := int64(250000 + rng.Intn(700001))
		status := statuses[rng.Intn(len(statuses))]
		condition := conditions[rng.Intn(len(conditions))]

		vehicles = append(vehicles, entity.Vehicle{
			PlateNumber: plate,
			Color:       color,
			Brand:       brand,
			Model:       model,
			CC:          cc,
			Year:        year,
			Mileage:     mileage,
			DailyRate:   dailyRate,
			Condition:   condition,
			Status:      status,
			Notes:       "seeded vehicle",
		})
	}

	if err := db.Create(&vehicles).Error; err != nil {
		return fmt.Errorf("create vehicles: %w", err)
	}

	return nil
}

func SeedRentals(db *gorm.DB) error {
	var count int64
	if err := db.Model(&entity.Rental{}).Count(&count).Error; err != nil {
		return fmt.Errorf("count rentals: %w", err)
	}

	if count >= seedRentalTarget {
		return nil
	}

	var customers []entity.Customer
	if err := db.Select("id").Find(&customers).Error; err != nil {
		return fmt.Errorf("query customers: %w", err)
	}
	if len(customers) == 0 {
		return fmt.Errorf("cannot seed rentals: no customers found")
	}

	var vehicles []entity.Vehicle
	if err := db.Select("id", "daily_rate", "mileage").Find(&vehicles).Error; err != nil {
		return fmt.Errorf("query vehicles: %w", err)
	}
	if len(vehicles) == 0 {
		return fmt.Errorf("cannot seed rentals: no vehicles found")
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	toCreate := seedRentalTarget - int(count)
	rentals := make([]entity.Rental, 0, toCreate)

	statuses := []entity.RentalStatus{
		entity.RentalStatusPending,
		entity.RentalStatusActive,
		entity.RentalStatusCompleted,
	}

	for i := 0; i < toCreate; i++ {
		customer := customers[rng.Intn(len(customers))]
		vehicle := vehicles[rng.Intn(len(vehicles))]

		totalDay := 1 + rng.Intn(14)
		startDate := time.Now().AddDate(0, 0, -(5 + rng.Intn(25)))
		endDate := startDate.AddDate(0, 0, totalDay)
		status := statuses[rng.Intn(len(statuses))]

		var returnDate *time.Time
		if status == entity.RentalStatusCompleted {
			r := endDate.AddDate(0, 0, rng.Intn(3))
			returnDate = &r
		}

		mileageStart := vehicle.Mileage
		mileageUsed := 20 + rng.Intn(981)
		mileageEnd := mileageStart + mileageUsed
		price := vehicle.DailyRate
		penaltyFee := int64(0)
		subtotal := price*int64(totalDay) + penaltyFee

		rentals = append(rentals, entity.Rental{
			CustomerID:            customer.ID,
			VehicleID:             vehicle.ID,
			StartDate:             startDate,
			EndDate:               endDate,
			TotalDay:              totalDay,
			ReturnDate:            returnDate,
			Price:                 price,
			PenaltyFee:            penaltyFee,
			Subtotal:              subtotal,
			Notes:                 "seeded rental",
			Status:                status,
			VehicleConditionStart: "Good",
			VehicleConditionEnd:   "Good",
			MileageStart:          mileageStart,
			MileageUsed:           mileageUsed,
			MileageEnd:            mileageEnd,
		})
	}

	if err := db.Create(&rentals).Error; err != nil {
		return fmt.Errorf("create rentals: %w", err)
	}

	return nil
}

func SeedVehicleIncidents(db *gorm.DB) error {
	var count int64
	if err := db.Model(&entity.VehicleIncident{}).Count(&count).Error; err != nil {
		return fmt.Errorf("count vehicle incidents: %w", err)
	}

	if count >= seedIncidentTarget {
		return nil
	}

	var rentals []entity.Rental
	if err := db.Select("id", "vehicle_id", "customer_id").Find(&rentals).Error; err != nil {
		return fmt.Errorf("query rentals: %w", err)
	}
	if len(rentals) == 0 {
		return fmt.Errorf("cannot seed vehicle incidents: no rentals found")
	}

	types := []constant.IncidentType{
		constant.IncidentAccident,
		constant.IncidentDamage,
		constant.IncidentTheft,
		constant.IncidentOther,
	}
	statuses := []constant.VehicleIncidentStatus{
		constant.VehicleIncidentStatusOpen,
		constant.VehicleIncidentStatusClosed,
	}
	descriptions := []string{
		"Front bumper scratched",
		"Minor side impact",
		"Rear light broken",
		"Tire damaged during trip",
		"Interior stain reported",
		"Mirror crack found",
		"Vehicle reported stolen and recovered",
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	toCreate := seedIncidentTarget - int(count)
	incidents := make([]entity.VehicleIncident, 0, toCreate)

	for i := 0; i < toCreate; i++ {
		rental := rentals[rng.Intn(len(rentals))]
		customerID := rental.CustomerID
		rentalID := rental.ID
		incidentType := types[rng.Intn(len(types))]
		status := statuses[rng.Intn(len(statuses))]
		description := descriptions[rng.Intn(len(descriptions))]

		penalty := int64(150000 + rng.Intn(2350001))
		if incidentType == constant.IncidentTheft {
			penalty = int64(5000000 + rng.Intn(5000001))
		}

		incidents = append(incidents, entity.VehicleIncident{
			VehicleID:    rental.VehicleID,
			CustomerID:   &customerID,
			RentalID:     &rentalID,
			IncidentDate: time.Now().AddDate(0, 0, -rng.Intn(90)),
			IncidentType: incidentType,
			Description:  description,
			PenaltyFee:   penalty,
			Status:       status,
		})
	}

	if err := db.Create(&incidents).Error; err != nil {
		return fmt.Errorf("create vehicle incidents: %w", err)
	}

	return nil
}
