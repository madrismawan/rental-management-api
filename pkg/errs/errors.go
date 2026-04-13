package errs

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDataConflict = errors.New("data conflict")
	ErrEmailDuplicate = errors.New("email already in use")

	ErrUserNotFound = errors.New("user not found")
	ErrDataNotFound = errors.New("data not found")
)