// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package sqlc

import (
	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}
