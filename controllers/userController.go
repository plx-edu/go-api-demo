package controllers

import (
	"encoding/json"
	"go-api-demo/custom"
	"go-api-demo/db"
	"go-api-demo/models"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	custom.Show("GetUsers()", 3)
	// w.Header().Set("Content-Type", "application/json")

	var users []models.User
	u := db.Instance.Find(&users)

	if u.Error != nil {
		log.Println("Error ", u.Error)

		err := map[string]interface{}{"status": 502, "message": "Cannot get users"}
		json.NewEncoder(w).Encode(err)
	}
	// fmt.Println(user)

	json.NewEncoder(w).Encode(users)
}

/*
func GetUser(w http.ResponseWriter, r *http.Request){
	custom.Show("GetUser()", 3)
	// w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Printf(":: Fetching user with id '%v' ::\n", params["id"])
	
	// var user models.User
	// fmt.Println(user)
	u := models.User{"John", 25}

	json.NewEncoder(w).Encode(u)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	custom.Show("CreateUser()", 3)
	
	// Assign user with POST data (r.Body)
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	
	if err != nil {
		custom.Show("error", 3)
	}

	// fmt.Println("::: ", err)
	// fmt.Println("::: ", user)

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	custom.Show("UpdateUser()", 3)

	params := mux.Vars(r)
	fmt.Printf(":: Fetching user with id '%v' ::\n", params["id"])
	
	// Assign user with POST data (r.Body)
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	
	if err != nil {
		custom.Show("error", 3)
	}

	// fmt.Println("::: ", err)
	// fmt.Println("::: ", user)

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	custom.Show("DeleteUser()", 3)
	
	params := mux.Vars(r)
	fmt.Printf(":: Fetching user with id '%v' ::\n", params["id"])
	
	// Assign user with POST data (r.Body)
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	
	if err != nil {
		custom.Show("error", 3)
	}

	// fmt.Println("::: ", err)
	// fmt.Println("::: ", user)

	json.NewEncoder(w).Encode(user)
}
*/