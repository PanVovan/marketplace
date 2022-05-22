package rating

import (
	"database/sql"
	"makretplace/internal/rating/presentation"
	"net/http"

	"github.com/gorilla/mux"
)

type Module struct {
}

func (m *Module) Configure(db *sql.DB, router *mux.Router) error {

	ratingRoute := router.PathPrefix("/rating").Subrouter()

	controller := presentation.NewRatingController()

	ratingRoute.HandleFunc("/product/{product_id}", controller.GetOne).Methods(http.MethodGet)
	ratingRoute.HandleFunc("/product/{product_id}/rate/{rate:[1-5]}", controller.Create).Methods(http.MethodPost)

	return nil
}
