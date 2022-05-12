package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-api-demo/controllers"
)


func Handler() *mux.Router{
	// init router
	r := mux.NewRouter().StrictSlash(true)
	r.Use(jsonHeader) // set default Content-Type to json



	// Route handles & endpoints
	r.HandleFunc("/", greetings).Methods("GET")
	r.HandleFunc("/about", about).Methods("GET")

	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	// r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	// r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	// r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	
	r.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	// r.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	// r.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
	// r.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT")
	// r.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", getHeaders)

	return r
}


// func getHeaders(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// }

func jsonHeader(next http.Handler)  http.Handler{
	return http.HandlerFunc(
		func (w http.ResponseWriter, r *http.Request) {
			// w.Header().Set("Content-Type", "application/json")
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
}


func greetings(w http.ResponseWriter, r *http.Request){
	// w.Write([]byte("Greetings & Salutations"))
	var g = map[string]interface{}{
		"int": 1,
		"string": "two",
		"boolean": true,
	}

	json.NewEncoder(w).Encode(g)
}

func about(w http.ResponseWriter, r *http.Request){
	// structs needs to be in capital letters
	type About struct {
		Who string
		Year int
	}

	about := About{"Some Dev", 2022}
	// b, _ := json.Marshal(about)

	json.NewEncoder(w).Encode(about)
}

// Personal notes:
// Gorm = ORM