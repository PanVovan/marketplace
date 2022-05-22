package user

import (
	"database/sql"
	"makretplace/internal/user/presentation"
	"net/http"

	"github.com/gorilla/mux"
)

type Module struct {
}

func (m *Module) Configure(db *sql.DB, router *mux.Router) error {

	userRoute := router.PathPrefix("/users").Subrouter()

	controller := presentation.NewUserController()

	userRoute.HandleFunc("/check", controller.Check).Methods(http.MethodGet)

	userRoute.HandleFunc("/login", controller.Login).Methods(http.MethodPost)

	userRoute.HandleFunc("/signup", controller.SignUp).Methods(http.MethodPost)

	return nil
}

// 	GET "/getall", userController.GetAll
// 	GET "/getone/:id", userController.GetOne
// 	POST "/create", userController.Create
// 	PUT "/update/:id", userController.Update
// 	DELETE "/delete/:id", userController.Delete
