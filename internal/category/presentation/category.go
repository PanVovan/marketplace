package presentation

import (
	"encoding/json"
	"makretplace/internal/category/application"
	response "makretplace/pkg/error"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CategoryController struct {
	getall *application.GetAllUseCase
	create *application.CreateUseCase
	getone *application.GetOneUseCase
	update *application.UpdateUseCase
	delete *application.DeleteUseCase
}

func NewCategoryController(
	getall *application.GetAllUseCase,
	create *application.CreateUseCase,
	getone *application.GetOneUseCase,
	update *application.UpdateUseCase,
	delete *application.DeleteUseCase,
) *CategoryController {
	return &CategoryController{
		getall: getall,
		create: create,
		update: update,
		delete: delete,
		getone: getone,
	}
}

func (cc *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
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

	categories, err := cc.getall.Execute(r.Context(), limit, page)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func (cc *CategoryController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	category, err := cc.getone.Execute(r.Context(), id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (cc *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	var create map[string]interface{}
	json.NewDecoder(r.Body).Decode(&create)

	id, err := cc.create.Execute(r.Context(), create["name"].(string))

	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := map[string]interface{}{
		"id": id,
	}

	json.NewEncoder(w).Encode(resp)
}

func (cc *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var update map[string]interface{}
	json.NewDecoder(r.Body).Decode(&update)

	err := cc.update.Execute(r.Context(), id, (update["name"].(string)))
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response.NewStatusResponse("ok"))
}

func (cc *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := cc.delete.Execute(r.Context(), id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response.NewStatusResponse("ok"))
}
