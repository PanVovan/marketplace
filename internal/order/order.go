package order

import (
	"database/sql"
	"makretplace/internal/order/presentation"
	"net/http"

	"github.com/gorilla/mux"
)

type Module struct {
}

func (m *Module) Configure(db *sql.DB, router *mux.Router) error {
	orderRoute := router.PathPrefix("/order").Subrouter()

	controller := presentation.NewOrderController()

	//admin
	orderRoute.HandleFunc("/admin/getall", controller.AdminDelete).Methods(http.MethodGet)

	orderRoute.HandleFunc("/admin/getone/{id}", controller.AdminGetOne).Methods(http.MethodGet)

	orderRoute.HandleFunc("/admin/getall/user/{id}", controller.AdminGetUser).Methods(http.MethodGet)

	orderRoute.HandleFunc("/admin/create", controller.AdminCreate).Methods(http.MethodPost)

	orderRoute.HandleFunc("/admin/delete/{id}", controller.AdminDelete).Methods(http.MethodDelete)

	//user
	orderRoute.HandleFunc("/user/getall", controller.UserGetAll).Methods(http.MethodGet)

	orderRoute.HandleFunc("/user/getone/{id}", controller.UserGetOne).Methods(http.MethodGet)

	orderRoute.HandleFunc("/user/create", controller.UserCreate).Methods(http.MethodPost)

	//guest
	orderRoute.HandleFunc("/guest/create", controller.GuestCreate).Methods(http.MethodPost)

	return nil
}
