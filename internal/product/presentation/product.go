package presentation

import (
	"encoding/json"
	"makretplace/internal/product/application"
	"makretplace/internal/product/domain/model"
	response "makretplace/pkg/error"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	create *application.CreateUseCase
	update *application.UpdateUseCase
	delete *application.DeleteUseCase
	getall *application.GetAllUseCase
	getone *application.GetOneUseCase
}

func NewProductController(
	create *application.CreateUseCase,
	getall *application.GetAllUseCase,
	getone *application.GetOneUseCase,
	update *application.UpdateUseCase,
	delete *application.DeleteUseCase,
) *ProductController {
	return &ProductController{
		create: create,
		update: update,
		delete: delete,
		getall: getall,
		getone: getone,
	}
}

func (pc *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	var (
		page       int       = 1
		limit      int       = 10
		err        error     = nil
		seller     *string   = nil
		brand      *string   = nil
		categories []*string = nil
	)

	if categoriesQuery, hasCategories := queries["categories"]; hasCategories {
		for _, category := range categoriesQuery {
			categories = append(categories, &category)
		}
	}

	if pageQuery, hasPage := queries["page"]; hasPage {
		if len(pageQuery) > 1 {
			response.NewErrorResponse(w, http.StatusInternalServerError, "too many page queries")
			return
		}
		page, err = strconv.Atoi(pageQuery[0])
		if err != nil {
			response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if limitQuery, hasLimit := queries["limit"]; hasLimit {
		if len(limitQuery) > 1 {
			response.NewErrorResponse(w, http.StatusInternalServerError, "too many limit queries")
			return
		}
		limit, err = strconv.Atoi(limitQuery[0])
		if err != nil {
			response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if sellerQuery, hasSeller := queries["seller"]; hasSeller {
		if len(sellerQuery) > 1 {
			response.NewErrorResponse(w, http.StatusInternalServerError, "too many seller queries")
			return
		}
		seller = &sellerQuery[0]
	}

	if brandQuery, hasBrand := queries["brand"]; hasBrand {
		if len(brandQuery) > 1 {
			response.NewErrorResponse(w, http.StatusInternalServerError, "too many seller queries")
			return
		}
		brand = &brandQuery[0]
	}

	products, err := pc.getall.Execute(r.Context(), seller, brand, categories, int32(limit), int32(page))

	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(products)
}

func (pc *ProductController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product model.Product

	product, err := pc.getone.Execute(r.Context(), id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (pc *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	var createProductParams model.CreateProductParams

	json.NewDecoder(r.Body).Decode(&createProductParams)
	id, err := pc.create.Execute(r.Context(), createProductParams)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := map[string]interface{}{
		"id": id,
	}

	json.NewEncoder(w).Encode(resp)

}

func (pc *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updateProductParams model.UpdateProductParams
	json.NewDecoder(r.Body).Decode(&updateProductParams)

	err := pc.update.Execute(r.Context(), id, updateProductParams)

	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response.NewStatusResponse("ok"))
}

func (pc *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := pc.delete.Execute(r.Context(), id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(response.NewStatusResponse("ok"))
}
