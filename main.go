package main

import (
	"labora-api/config"
	"labora-api/controllers"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	config.InitSetUp()
	router := mux.NewRouter()

	router.HandleFunc("/items", controllers.GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", controllers.BuscarID).Methods("GET")
	router.HandleFunc("/items", controllers.CrearItem).Methods("POST")
	router.HandleFunc("/items/{id}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")
	router.HandleFunc("/items/search/name", controllers.GetItemByName).Methods("GET")

	port := ":9000"
	if err := config.StartServer(port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
