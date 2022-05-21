package dao

import (
	"database/sql"
	"makretplace/internal/rating/infrastructure/database/sqlc"
)

type RatingDao struct {
	db *sql.DB
	*sqlc.Queries
}

func NewRatingDao(db *sql.DB) *RatingDao {
	return &RatingDao{
		db:      db,
		Queries: sqlc.New(db),
	}
}
