package presentation

import (
	"encoding/json"
	"net/http"
)

type SellerController struct {
}

func NewSellerController() *SellerController {
	return &SellerController{}
}

func (sc *SellerController) Login(w http.ResponseWriter, r *http.Request) {
	var name map[string]interface{}
	json.NewDecoder(r.Body).Decode(&name)

	resp := map[string]interface{}{
		"login": "login " + name["name"].(string),
	}

	json.NewEncoder(w).Encode(resp)

}

func (sc *SellerController) Check(w http.ResponseWriter, r *http.Request) {

	var name map[string]interface{}
	json.NewDecoder(r.Body).Decode(&name)

	resp := map[string]interface{}{
		"message": "check " + name["name"].(string),
	}

	json.NewEncoder(w).Encode(resp)
}

func (sc *SellerController) SignUp(w http.ResponseWriter, r *http.Request) {

	var name map[string]interface{}
	json.NewDecoder(r.Body).Decode(&name)

	resp := map[string]interface{}{
		"message": "signup " + name["name"].(string),
	}

	json.NewEncoder(w).Encode(resp)
}
