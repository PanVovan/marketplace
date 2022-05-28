package presentation

import (
	"encoding/json"
	"makretplace/internal/product/application"
	"makretplace/internal/product/domain/model"
	response "makretplace/pkg/error"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductPropertyController struct {
	create *application.CreatePropertyUseCase
	update *application.UpdatePropertyUseCase
	delete *application.DeletePropertyUseCase
	getone *application.GetOnePropertyUseCase
	getall *application.GetAllPropertyUseCase
}

func NewProductPropertyController(
	create *application.CreatePropertyUseCase,
	getall *application.GetAllPropertyUseCase,
	getone *application.GetOnePropertyUseCase,
	update *application.UpdatePropertyUseCase,
	delete *application.DeletePropertyUseCase,
) *ProductPropertyController {
	return &ProductPropertyController{
		create: create,
		update: update,
		delete: delete,
		getall: getall,
		getone: getone,
	}
}

func (pc *ProductPropertyController) GetAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["product_id"]
	properties, err := pc.getall.Execute(r.Context(), productId)

	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(properties)
}

func (pc *ProductPropertyController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//productId := vars["product_id"]

	var property model.ProductProperty

	property, err := pc.getone.Execute(r.Context(), id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(property)
}

func (pc *ProductPropertyController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["product_id"]

	var create model.CreateProductPropertyParams

	json.NewDecoder(r.Body).Decode(&create)

	id, err := pc.create.Execute(r.Context(), productId, create)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := map[string]interface{}{
		"id": id,
	}

	json.NewEncoder(w).Encode(resp)
}

func (pc *ProductPropertyController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//productId := vars["product_id"]

	var update model.UpdateProductPropertyParams
	json.NewDecoder(r.Body).Decode(&update)

	err := pc.update.Execute(r.Context(), id, update)

	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response.NewStatusResponse("ok"))
}

func (pc *ProductPropertyController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//productId := vars["product_id"]

	err := pc.delete.Execute(r.Context(), id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response.NewStatusResponse("ok"))
}
