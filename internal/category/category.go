package category

import (
	"database/sql"
	"makretplace/internal/category/presentation"
	"net/http"

	"github.com/gorilla/mux"
)

type Module struct {
}

func (m *Module) Configure(db *sql.DB, router *mux.Router) error {
	categoryRoute := router.PathPrefix("/category").Subrouter()

	controller := presentation.NewCategoryController()

	categoryRoute.HandleFunc("/getall", controller.GetAll).Methods(http.MethodGet)

	categoryRoute.HandleFunc("/getone/{id}", controller.GetOne).Methods(http.MethodGet)

	categoryRoute.HandleFunc("/create", controller.Create).Methods(http.MethodPost)

	categoryRoute.HandleFunc("/update/{id}", controller.Update).Methods(http.MethodPut)

	categoryRoute.HandleFunc("/delete/{id}", controller.Delete).Methods(http.MethodDelete)

	return nil
}
