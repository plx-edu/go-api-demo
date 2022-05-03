package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const JSON = "application/json"

func Handler() *mux.Router{
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", greetings).Methods("GET")
	r.HandleFunc("/about", about).Methods("GET")

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveInit)
	return r
}

func serveInit(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", JSON)
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func greetings(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Greetings & Salutations"))
}

func about(w http.ResponseWriter, r *http.Request){
	// structs needs to be in capital letters
	type User struct {
		Name string
		Age int
	}

	u := User{"John", 25}
	b, _ := json.Marshal(u)


	fmt.Println("a: ",u)
	fmt.Println("b: ", string(b))

	json.NewEncoder(w).Encode(u)
}