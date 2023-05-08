package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

//psql -h localhost -p 5431 -U alfred -d items

const (
	host        = "localhost"
	port        = "5431"
	dbName      = "items"
	rolName     = "alfred"
	rolPassword = "4lfr3d"
)

func Connect_BD() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful connection to the database:", dbConn)
	return dbConn, err
}

func StartServer(port string, router http.Handler) error {
	servidor := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting alfred on port %s...\n", port)
	if err := servidor.ListenAndServe(); err != nil {
		return fmt.Errorf("Error while starting up alfred: '%v'", err)
	}

	return nil
}

func InitSetUp() {
	db, err := Connect_BD()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
