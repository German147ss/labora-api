package services

import (
	"database/sql"
	"fmt"
	"log"
)

type DbConnection struct {
	*sql.DB
}

var Db DbConnection

func UpDb() {
	err := Connect_BD()
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DbConnection) PingOrDie() {
	if err := db.Ping(); err != nil {
		log.Fatalf("can't reach databse, error: %v", err)
	}
}

//psql -h localhost -p 5431 -U alfred -d items

const (
	host        = "localhost"
	port        = "5431"
	dbName      = "items"
	rolName     = "alfred"
	rolPassword = "4lfr3d"
)

var dbConn *sql.DB

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
