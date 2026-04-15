package constant

type CustomerLogStatus string

const (
	CustomerLogStatusActive   CustomerLogStatus = "active"
	CustomerLogStatusInactive CustomerLogStatus = "inactive"
	CustomerLogStatusBanned   CustomerLogStatus = "banned"
)
