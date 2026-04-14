package constant

type RentalStatus string

const (
	RentalStatusPending   RentalStatus = "pending"
	RentalStatusActive    RentalStatus = "active"
	RentalStatusCompleted RentalStatus = "completed"
	RentalStatusCancelled RentalStatus = "cancelled"
)
