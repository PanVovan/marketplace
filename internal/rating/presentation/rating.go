package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RatingController struct {
}

func NewRatingController() *RatingController {
	return &RatingController{}
}

func (rc *RatingController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["product_id"]
	w.Write([]byte("Getting rating of " + productId))
}

func (rc *RatingController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rate := vars["rate"]
	productId := vars["product_id"]
	w.Write([]byte("Creating rating " + rate + " for product " + productId))

}
