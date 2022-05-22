package basket

import (
	"database/sql"
	"makretplace/internal/basket/presentation"
	"net/http"

	"github.com/gorilla/mux"
)

type Module struct {
}

func (m *Module) Configure(db *sql.DB, router *mux.Router) error {
	basketRoute := router.PathPrefix("/basket").Subrouter()

	controller := presentation.NewBasketController()

	basketRoute.HandleFunc("/getone", controller.GetOne).Methods(http.MethodGet)

	basketRoute.HandleFunc("/product/{product_id}/append/{quantity:[0-9]+}", controller.Append).Methods(http.MethodPut)

	basketRoute.HandleFunc("/product/{product_id}/increment/{quantity:[0-9]+}", controller.Increment).Methods(http.MethodPut)

	basketRoute.HandleFunc("/product/{product_id}/decrement/{quantity:[0-9]+}", controller.Decrement).Methods(http.MethodPut)

	basketRoute.HandleFunc("/product/{product_id}/remove", controller.Remove).Methods(http.MethodPut)

	basketRoute.HandleFunc("/clear", controller.Clear).Methods(http.MethodPut)

	return nil
}
