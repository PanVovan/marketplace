package product

import (
	"database/sql"
	"makretplace/internal/product/presentation"
	"net/http"

	"github.com/gorilla/mux"
)

type Module struct {
}

func (m *Module) Configure(db *sql.DB, router *mux.Router) error {
	productRoute := router.PathPrefix("/product").Subrouter()

	productController := presentation.NewProductController()
	productPropertyController := presentation.NewProductPropertyController()

	//Products
	productRoute.HandleFunc("/getall", productController.GetAll).Methods(http.MethodGet)

	productRoute.HandleFunc("/getone/{id}", productController.GetOne).Methods(http.MethodGet)

	productRoute.HandleFunc("/create", productController.Create).Methods(http.MethodPost)

	productRoute.HandleFunc("/update/{id}", productController.Update).Methods(http.MethodPut)

	productRoute.HandleFunc("/delete/{id}", productController.Delete).Methods(http.MethodDelete)

	//Properties
	productRoute.HandleFunc("/{product_id}/property/getall", productPropertyController.GetAll).Methods(http.MethodGet)

	productRoute.HandleFunc("/{product_id}/property/getone/{id}", productPropertyController.GetOne).Methods(http.MethodGet)

	productRoute.HandleFunc("/{product_id}/property/create", productPropertyController.Create).Methods(http.MethodPost)

	productRoute.HandleFunc("/{product_id}/property/update/{id}", productPropertyController.Update).Methods(http.MethodPut)

	productRoute.HandleFunc("/{product_id}/property/delete/{id}", productPropertyController.Delete).Methods(http.MethodDelete)

	return nil
}
