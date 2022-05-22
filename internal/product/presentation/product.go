package presentation

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type ProductController struct {
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (pc *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	categories := queries["categories"]
	brand := queries["brand"]
	seller := queries["seller"]

	w.Write([]byte("Getting ALL " +
		" where seller: " + strings.Join(seller, "") +
		" brand: " + strings.Join(brand, "") +
		" categories: " + strings.Join(categories, ", ")))
}

func (pc *ProductController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Write([]byte("Getting product " + id))
}

func (pc *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created"))
}

func (pc *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Updated " + id))
}

func (pc *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Deleted " + id))
}
