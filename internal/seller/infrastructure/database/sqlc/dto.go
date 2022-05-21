package sqlc

import (
	"github.com/google/uuid"
)

type SellerInfo struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}
