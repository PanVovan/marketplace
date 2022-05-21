package presentation

import (
	"encoding/json"
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var name map[string]interface{}
	json.NewDecoder(r.Body).Decode(&name)

	resp := map[string]interface{}{
		"login": "login " + name["name"].(string),
	}

	json.NewEncoder(w).Encode(resp)

}

func (uc *UserController) Check(w http.ResponseWriter, r *http.Request) {

	var name map[string]interface{}
	json.NewDecoder(r.Body).Decode(&name)

	resp := map[string]interface{}{
		"message": "check " + name["name"].(string),
	}

	json.NewEncoder(w).Encode(resp)
}

func (uc *UserController) SignUp(w http.ResponseWriter, r *http.Request) {

	var name map[string]interface{}
	json.NewDecoder(r.Body).Decode(&name)

	resp := map[string]interface{}{
		"message": "signup " + name["name"].(string),
	}

	json.NewEncoder(w).Encode(resp)
}
