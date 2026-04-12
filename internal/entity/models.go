package entity

// Models returns all entities for AutoMigrate.
func Models() []any {
	return []any{
		&User{},
		&Customer{},
		&Vehicle{},
		&Rental{},
		&VehicleIncident{},
	}
}
