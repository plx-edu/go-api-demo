package controllers

import (
	"encoding/json"
	"go-api-demo/custom"
	"go-api-demo/models"
	"go-api-demo/services"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	custom.AppLog(r)

	json.NewEncoder(w).Encode(services.GetUsers())
}


func CreateUser(w http.ResponseWriter, r *http.Request){
	custom.AppLog(r)

	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	services.CreateUser(*user)

	if user.ID == 0 {
		json.NewEncoder(w).Encode(custom.GenericError())
		return
	}

	json.NewEncoder(w).Encode(user)
}
