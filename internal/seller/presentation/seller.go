package presentation

import (
	"encoding/json"
	"makretplace/internal/seller/application"
	"makretplace/internal/seller/domain/model"
	error_response "makretplace/pkg/error"
	"net/http"
)

type SellerController struct {
	signup *application.SignUpUseCase
	login  *application.LoginUseCase
}

func NewSellerController(
	signup *application.SignUpUseCase,
	login *application.LoginUseCase,
) *SellerController {
	return &SellerController{
		signup: signup,
		login:  login,
	}
}

func (uc *SellerController) Login(w http.ResponseWriter, r *http.Request) {
	var seller application.LoginSellerParams
	json.NewDecoder(r.Body).Decode(&seller)

	token, err := uc.login.Execute(r.Context(), seller)

	if err != nil {
		error_response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	resp := map[string]interface{}{
		"token": token,
	}

	json.NewEncoder(w).Encode(resp)

}

func (uc *SellerController) Check(w http.ResponseWriter, r *http.Request) {

	var name map[string]interface{}
	json.NewDecoder(r.Body).Decode(&name)

	resp := map[string]interface{}{
		"token": "check " + name["name"].(string),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (uc *SellerController) SignUp(w http.ResponseWriter, r *http.Request) {

	var sellerParams model.CreateSellerParams

	json.NewDecoder(r.Body).Decode(&sellerParams)

	token, id, err := uc.signup.Execute(r.Context(), sellerParams)

	if err != nil {
		error_response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := map[string]interface{}{
		"id":    id,
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
