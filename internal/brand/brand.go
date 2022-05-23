package brand

import (
	"database/sql"
	"makretplace/internal/brand/application"
	"makretplace/internal/brand/infrastructure/database"
	"makretplace/internal/brand/presentation"
	"net/http"

	"github.com/gorilla/mux"
)

type Module struct {
}

func (m *Module) Configure(db *sql.DB, router *mux.Router) error {
	brandRoute := router.PathPrefix("/brand").Subrouter()

	repository := database.NewBrandRepositoryPostgres(db)

	getall := application.NewGetAllUseCase(repository)
	create := application.NewCreateUseCase(repository)
	getone := application.NewGetOneUseCase(repository)
	update := application.NewUpdateUseCase(repository)
	delete := application.NewDeleteUseCase(repository)

	controller := presentation.NewBrandController(
		getall,
		create,
		getone,
		update,
		delete,
	)

	brandRoute.HandleFunc("/getall", controller.GetAll).Methods(http.MethodGet)

	brandRoute.HandleFunc("/getone/{id}", controller.GetOne).Methods(http.MethodGet)

	brandRoute.HandleFunc("/create", controller.Create).Methods(http.MethodPost)

	brandRoute.HandleFunc("/update/{id}", controller.Update).Methods(http.MethodPut)

	brandRoute.HandleFunc("/delete/{id}", controller.Delete).Methods(http.MethodDelete)

	return nil
}
