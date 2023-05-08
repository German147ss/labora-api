package services

import (
	"database/sql"
	"fmt"
	"log"
)

// DbConnection contains a pointer to the SQL database.
type DbConnection struct {
	*sql.DB
}

var Db DbConnection

// UpDb connects to the database.
func UpDb() {
	err := Connect_BD()
	if err != nil {
		log.Fatal(err)
	}
}

// PingOrDie pings the database and logs a fatal error if it can't be reached.
func (db *DbConnection) PingOrDie() {
	if err := db.Ping(); err != nil {
		log.Fatalf("can't reach database, error: %v", err)
	}
}

// Constants used to connect to the database.
const (
	host        = "localhost"
	port        = "5431"
	dbName      = "items"
	rolName     = "alfred"
	rolPassword = "4lfr3d"
)

var dbConn *sql.DB

// Connect_BD connects to the database and returns an error if the connection fails.
func Connect_BD() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	var err error
	dbConn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful connection to the database:", dbConn)
	Db = DbConnection{dbConn}
	Db.PingOrDie()
	return err
}
