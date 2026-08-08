package user

import (
	"github.com/google/uuid"
)

// User represents an individual user entity.
type User struct {
	ID                 uuid.UUID `json:"id"`
	CreatedAtTimestamp int64     `json:"created_at_timestamp"`
	Username           string    `json:"username"`
	EmailVerified      bool      `json:"email_verified"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	PrimaryEmail       string    `json:"primary_email"`
	SecondaryEmail     string    `json:"secondary_email"`
}
