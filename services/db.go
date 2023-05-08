package services

import (
	"database/sql"
	"fmt"
	"log"
)

// DbConnection contiene un puntero a la base de datos SQL.
type DbConnection struct {
	*sql.DB
}

var Db DbConnection

// UpDb conecta con la base de datos.
func UpDb() {
	err := Connect_BD()
	if err != nil {
		log.Fatal(err)
	}
}

// PingOrDie envía un ping a la base de datos y si no se puede alcanzar, registra un error fatal.
func (db *DbConnection) PingOrDie() {
	if err := db.Ping(); err != nil {
		log.Fatalf("no se puede alcanzar la base de datos, error: %v", err)
	}
}

// Constantes utilizadas para conectarse a la base de datos.
const (
	host        = "localhost"
	port        = "5431"
	dbName      = "items"
	rolName     = "alfred"
	rolPassword = "4lfr3d"
)

var dbConn *sql.DB

// Connect_BD conecta con la base de datos y devuelve un error si falla la conexión.
func Connect_BD() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	var err error
	dbConn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexión exitosa a la base de datos:", dbConn)
	Db = DbConnection{dbConn}
	Db.PingOrDie()
	return err
}
