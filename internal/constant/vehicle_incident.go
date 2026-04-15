package constant

type VehicleIncidentStatus string
type IncidentType string

const (
	VehicleIncidentStatusOpen       VehicleIncidentStatus = "open"
	VehicleIncidentStatusInProgress VehicleIncidentStatus = "in_progress"
	VehicleIncidentStatusResolved   VehicleIncidentStatus = "resolved"
	VehicleIncidentStatusClosed     VehicleIncidentStatus = "closed"
)

const (
	IncidentAccident IncidentType = "accident"
	IncidentDamage   IncidentType = "damage"
	IncidentTheft    IncidentType = "theft"
	IncidentOther    IncidentType = "other"
)
