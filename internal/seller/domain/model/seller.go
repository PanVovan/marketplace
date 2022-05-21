package model

import "github.com/google/uuid"

type Seller struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
}
