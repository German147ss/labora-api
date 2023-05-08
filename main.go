package main

import (
	"labora-api/config"
	"labora-api/controllers"
	"labora-api/services"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	services.UpDb()
	router := mux.NewRouter()

	router.HandleFunc("/items", controllers.GetItemsHandler).Methods("GET")
	router.HandleFunc("/items/simple", controllers.ObtenerItems).Methods("GET")

	router.HandleFunc("/items/{id}", controllers.BuscarID).Methods("GET")
	router.HandleFunc("/items", controllers.CrearItem).Methods("POST")
	router.HandleFunc("/items/{id}", controllers.EditarItemHandler).Methods("PUT")
	router.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")
	router.HandleFunc("/items/search/name", controllers.GetItemByName).Methods("GET")

	services.Db.PingOrDie()
	port := ":9000"
	if err := config.StartServer(port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
