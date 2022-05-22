package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ProductPropertyController struct {
}

func NewProductPropertyController() *ProductPropertyController {
	return &ProductPropertyController{}
}

func (pc *ProductPropertyController) GetAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["product_id"]
	w.Write([]byte("Getting ALL property from product " + productId))
}

func (pc *ProductPropertyController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productId := vars["product_id"]
	w.Write([]byte("Getting property " + id + " from product " + productId))
}

func (pc *ProductPropertyController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["product_id"]
	w.Write([]byte("Creating property to product " + productId))
}

func (pc *ProductPropertyController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productId := vars["product_id"]
	w.Write([]byte("Updating property " + id + " from product " + productId))
}

func (pc *ProductPropertyController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productId := vars["product_id"]
	w.Write([]byte("Deleting property " + id + " from product " + productId))
}
