package presentation

import (
	"encoding/json"
	"fmt"
	"makretplace/internal/user/application"
	"makretplace/internal/user/domain/model"
	"net/http"
)

type UserController struct {
	signup *application.SignUpUseCase
	login  *application.LoginUseCase
}

func NewUserController(
	signup *application.SignUpUseCase,
	login *application.LoginUseCase,
) *UserController {
	return &UserController{
		signup: signup,
		login:  login,
	}
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user application.LoginUserParams
	json.NewDecoder(r.Body).Decode(&user)

	token, _ := uc.login.Execute(r.Context(), user)

	resp := map[string]interface{}{
		"token": token,
	}

	json.NewEncoder(w).Encode(resp)

}

func (uc *UserController) Check(w http.ResponseWriter, r *http.Request) {

	var name map[string]interface{}
	json.NewDecoder(r.Body).Decode(&name)

	resp := map[string]interface{}{
		"token": "check " + name["name"].(string),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (uc *UserController) SignUp(w http.ResponseWriter, r *http.Request) {

	var userParams model.CreateUserParams

	json.NewDecoder(r.Body).Decode(&userParams)

	token, id, err := uc.signup.Execute(r.Context(), userParams)

	if err != nil {
		fmt.Println(err)
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
