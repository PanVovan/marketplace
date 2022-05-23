package presentation

import (
	"encoding/json"
	"net/http"
	"strconv"

	"makretplace/internal/brand/application"
	response "makretplace/pkg/error"

	"github.com/gorilla/mux"
)

type BrandController struct {
	getall *application.GetAllUseCase
	create *application.CreateUseCase
	getone *application.GetOneUseCase
	update *application.UpdateUseCase
	delete *application.DeleteUseCase
}

func NewBrandController(
	getall *application.GetAllUseCase,
	create *application.CreateUseCase,
	getone *application.GetOneUseCase,
	update *application.UpdateUseCase,
	delete *application.DeleteUseCase,
) *BrandController {
	return &BrandController{
		getall: getall,
		create: create,
		update: update,
		delete: delete,
		getone: getone,
	}
}

func (bc *BrandController) GetAll(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	var page int32 = 1
	var limit int32 = 10

	pageQuery, haspage := queries["page"]
	limitQuery, haslimit := queries["limit"]

	if haspage {
		if len(pageQuery) > 1 {
			response.NewErrorResponse(w, http.StatusBadRequest, "Too many page parametrs")
			return
		}
		page_q, err := strconv.Atoi(pageQuery[0])
		if err != nil {
			response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		page = int32(page_q)
	}
	if haslimit {
		if len(limitQuery) > 1 {
			response.NewErrorResponse(w, http.StatusBadRequest, "Too many limit parametrs")
			return
		}

		limit_q, err := strconv.Atoi(limitQuery[0])
		if err != nil {
			response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		limit = int32(limit_q)
	}

	brands, err := bc.getall.Execute(r.Context(), limit, page)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(brands)
}

func (bc *BrandController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	brand, err := bc.getone.Execute(r.Context(), id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(brand)
}

func (bc *BrandController) Create(w http.ResponseWriter, r *http.Request) {
	var create map[string]interface{}
	json.NewDecoder(r.Body).Decode(&create)

	id, err := bc.create.Execute(r.Context(), create["name"].(string))

	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := map[string]interface{}{
		"id": id,
	}

	json.NewEncoder(w).Encode(resp)

}

func (bc *BrandController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var update map[string]interface{}
	json.NewDecoder(r.Body).Decode(&update)

	err := bc.update.Execute(r.Context(), id, (update["name"].(string)))
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response.NewStatusResponse("ok"))
}

func (bc *BrandController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := bc.delete.Execute(r.Context(), id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response.NewStatusResponse("ok"))
}
