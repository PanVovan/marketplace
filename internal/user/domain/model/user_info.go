package model

import "github.com/google/uuid"

type UserInfo struct {
	ID   uuid.UUID
	Name string
}
