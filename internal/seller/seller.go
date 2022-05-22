package seller

import (
	"database/sql"
	"makretplace/internal/seller/application"
	"makretplace/internal/seller/infrastructure/database"
	"makretplace/internal/seller/infrastructure/service"
	"makretplace/internal/seller/presentation"
	"net/http"

	"github.com/gorilla/mux"
)

type Module struct {
}

func (m *Module) Configure(db *sql.DB, router *mux.Router) error {

	sellerRoute := router.PathPrefix("/seller").Subrouter()

	repository := database.NewSellerRepositoryPostgres(db)
	userService := service.NewSellerService()
	signupUseCase := application.NewSignUpUseCase(repository, userService)
	loginUseCase := application.NewLoginUseCase(repository, userService)
	controller := presentation.NewSellerController(signupUseCase, loginUseCase)

	sellerRoute.HandleFunc("/check", controller.Check).Methods(http.MethodGet)

	sellerRoute.HandleFunc("/login", controller.Login).Methods(http.MethodPost)

	sellerRoute.HandleFunc("/signup", controller.SignUp).Methods(http.MethodPost)

	return nil
}
