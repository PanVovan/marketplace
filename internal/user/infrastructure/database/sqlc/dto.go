package sqlc

import (
	"github.com/google/uuid"
)

type UserInfo struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}
