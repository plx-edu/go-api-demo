package main

import (
	"fmt"
	"go-api-demo/custom"
	"go-api-demo/db"
	"go-api-demo/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Application Entrypoint
func main(){
	env := godotenv.Load()

	if env  != nil {
		log.Fatal("Error loading .env file")
	}
	
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	
	db.Connect()
	db.Migrate()
	
	// Start server
	custom.Show(fmt.Sprintf("Listening on port %v", port), 5)
	log.Fatal(http.ListenAndServe(port, routes.Handler()))
}
