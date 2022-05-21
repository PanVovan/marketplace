package model

import "github.com/google/uuid"

type ProductProperty struct {
	ID    uuid.UUID
	Name  string
	Value string
}
