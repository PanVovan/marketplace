package entity

import "github.com/google/uuid"

// User model, DB table: "users"
type User struct {
	ID       uuid.UUID `json:"-" db:"id"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	Role     string    `json:"role" db:"role"`
}
