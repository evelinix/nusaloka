package handler

import (
	"encoding/json"
	"net/http"

	"github.com/evelinix/nusaloka/internal/account/dto"
	"github.com/evelinix/nusaloka/internal/account/model"
	"github.com/evelinix/nusaloka/internal/account/service"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user dto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var userData = model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	err = service.RegisterUser(userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}