package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryController struct {
}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

func (cc *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting ALL"))
}

func (cc *CategoryController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Getting " + id))
}

func (cc *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created"))
}

func (cc *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Updated " + id))
}

func (cc *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte("Deleted " + id))
}
