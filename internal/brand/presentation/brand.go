package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
)

type BrandController struct {
}

func NewBrandController() *BrandController {
	return &BrandController{}
}

func (bc *BrandController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting ALL"))
}

func (bc *BrandController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Getting " + id))
}

func (bc *BrandController) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created"))
}

func (bc *BrandController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Updated " + id))
}

func (bc *BrandController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Deleted " + id))
}
