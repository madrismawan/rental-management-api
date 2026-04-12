package constant

type VehicleStatus string

const (
	VehicleStatusAvailable   VehicleStatus = "available"
	VehicleStatusRented      VehicleStatus = "rented"
	VehicleStatusMaintenance VehicleStatus = "maintenance"
)
