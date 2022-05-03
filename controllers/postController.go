package controllers

import (
	"encoding/json"
	"go-api-demo/custom"
	"go-api-demo/models"
	"net/http"
)


func GetPosts(w http.ResponseWriter, r *http.Request){
	custom.Show("GetPosts()", 3)
	
	var posts []models.Post
	// posts := models.Post{"Title", "Text Content", &models.User{ "Lunah", 20}}

	json.NewEncoder(w).Encode(posts)
}

/*
func GetPost(w http.ResponseWriter, r *http.Request){
	custom.Show("GetPost()", 3)

	
	params := mux.Vars(r)
	fmt.Printf(":: Fetching post with id '%v' ::\n", params["id"])
	
	posts := models.Post{"Title", "Text Content", &models.User{ "Lunah", 20}}

	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request){
	custom.Show("CreatePost()", 3)
	
	post := &models.Post{}
	err := json.NewDecoder(r.Body).Decode(post)

	if err != nil {
		custom.Show("Error creating Post", 3)
	}

	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request){
	custom.Show("UpdatePost()", 3)

	
	params := mux.Vars(r)
	fmt.Printf(":: Fetching post with id '%v' ::\n", params["id"])
	
	posts := models.Post{"Title", "Text Content", &models.User{ "Lunah", 20}}

	json.NewEncoder(w).Encode(posts)
}

func DeletePost(w http.ResponseWriter, r *http.Request){
	custom.Show("DeletePost()", 3)

	
	params := mux.Vars(r)
	fmt.Printf(":: Fetching post with id '%v' ::\n", params["id"])
	
	posts := models.Post{"Title", "Text Content", &models.User{ "Lunah", 20}}

	json.NewEncoder(w).Encode(posts)
}
*/