package user

import (
	"database/sql"
	"makretplace/internal/user/application"
	"makretplace/internal/user/infrastructure/database"
	"makretplace/internal/user/infrastructure/service"
	"makretplace/internal/user/presentation"
	"net/http"

	"github.com/gorilla/mux"
)

type Module struct {
}

func (m *Module) Configure(db *sql.DB, router *mux.Router) error {

	userRoute := router.PathPrefix("/user").Subrouter()

	repository := database.NewUserRepositoryPostgres(db)
	userService := service.NewUserService()
	signupUseCase := application.NewSignUpUseCase(repository, userService)
	loginUseCase := application.NewLoginUseCase(repository, userService)
	controller := presentation.NewUserController(signupUseCase, loginUseCase)

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
