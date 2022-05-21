package model

import "github.com/google/uuid"

type ProductImage struct {
	ID   uuid.UUID
	File string
}
