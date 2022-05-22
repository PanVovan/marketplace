package presentation

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BasketController struct {
}

func NewBasketController() *BasketController {
	return &BasketController{}
}

func (bc *BasketController) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting One"))
}

func (bc *BasketController) Append(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["product_id"]

	quantity, err := strconv.Atoi(vars["quantity"])

	if err != nil {
		return
	}

	w.Write([]byte("Append product " + productId + " to basket in a quantity of " + strconv.Itoa(quantity)))
}

func (bc *BasketController) Increment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["product_id"]

	quantity, err := strconv.Atoi(vars["quantity"])

	if err != nil {
		return
	}

	w.Write([]byte("Increment product " + productId + " to basket by " + strconv.Itoa(quantity)))
}

func (bc *BasketController) Decrement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["product_id"]

	quantity, err := strconv.Atoi(vars["quantity"])

	if err != nil {
		return
	}

	w.Write([]byte("Increment product " + productId + " to basket by " + strconv.Itoa(quantity)))
}

func (bc *BasketController) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["product_id"]

	w.Write([]byte("Remove product " + productId))
}

func (bc *BasketController) Clear(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Clear basket"))
}

// router.get('/getone', BasketController.getOne)
// router.put('/product/:productId([0-9]+)/append/:quantity([0-9]+)', BasketController.append)
// router.put('/product/:productId([0-9]+)/increment/:quantity([0-9]+)', BasketController.increment)
// router.put('/product/:productId([0-9]+)/decrement/:quantity([0-9]+)', BasketController.decrement)
// router.put('/product/:productId([0-9]+)/remove', BasketController.remove)
// router.put('/clear', BasketController.clear)
