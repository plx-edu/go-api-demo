package main

import (
	"fmt"
	"go-api-demo/routes"
	"net/http"
)

func main(){
	// Simple Hello World
	message := decorate("Greetings and Salutations", 5)
	fmt.Println(message)

	// Using routes
	// http.Handle("/", routes.Handler())

	http.ListenAndServe(":3001", routes.Handler())

}

func decorate(msg string, n int) string{
	dots := ""

	for i := 0; i < n; i++ {
		dots += ":"
	}

	return dots+" "+msg+" "+dots
}