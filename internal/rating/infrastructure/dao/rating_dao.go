package dao

import (
	"database/sql"
	"main/internal/rating/infrastructure/sqlc"
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
