package constant

type VehicleStatus string
type VehicleCondition string

const (
	VehicleStatusAvailable   VehicleStatus    = "available"
	VehicleStatusRented      VehicleStatus    = "rented"
	VehicleStatusMaintenance VehicleStatus    = "maintenance"
	VehicleStatusUnavailable VehicleStatus    = "unavailable"
	VehicleConditionGood     VehicleCondition = "good"
	VehicleConditionBroke    VehicleCondition = "broke"
	VehicleConditionService  VehicleCondition = "service"
)
