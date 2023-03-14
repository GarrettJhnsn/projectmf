package main

import (
	"log"
	"net/http"

	"github.com/garrettjhnsn/projectmf/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("Starting application on port:", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))

}
