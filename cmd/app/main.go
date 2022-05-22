package main

import (
	"makretplace/internal/basket"
	"makretplace/internal/brand"
	"makretplace/internal/category"
	"makretplace/internal/order"
	"makretplace/internal/product"
	"makretplace/internal/rating"
	"makretplace/internal/seller"
	"makretplace/internal/user"
	marketplace "makretplace/pkg/httpserver"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	routes := mux.NewRouter()

	userModule := &user.Module{}
	userModule.Configure(nil, routes)

	sellerModule := &seller.Module{}
	sellerModule.Configure(nil, routes)

	productModule := &product.Module{}
	productModule.Configure(nil, routes)

	categoryModule := &category.Module{}
	categoryModule.Configure(nil, routes)

	brandModule := &brand.Module{}
	brandModule.Configure(nil, routes)

	basketModule := &basket.Module{}
	basketModule.Configure(nil, routes)

	orderModule := &order.Module{}
	orderModule.Configure(nil, routes)

	ratingModule := &rating.Module{}
	ratingModule.Configure(nil, routes)

	server := new(marketplace.Server)
	server.Run("localhost", "8080", routes)
}
