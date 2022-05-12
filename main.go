package main

import (
	"fmt"
	"go-api-demo/custom"
	"go-api-demo/db"
	"go-api-demo/routes"
	"go-api-demo/services"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Application Entrypoint
func main(){
	// Load env vars from a .env file
	env := godotenv.Load()

	if env  != nil {
		log.Fatal("::::: Error loading .env file")
	}
	
	// For a successful migration
	// Capitalize fields in models
	db.Connect()
	db.Migrate()
	
	// Retrieve port number from env
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	fmt.Printf("Users: %v\n",services.GetUsers())
	
	// Start server
	custom.Show(fmt.Sprintf("Listening on port %v", port), 5)
	log.Fatal(http.ListenAndServe(port, routes.Handler()))
}
