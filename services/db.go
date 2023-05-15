package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DbConnection contiene un puntero a la base de datos SQL.
type DbConnection struct {
	*sql.DB
}

var DataBasePAPA DbConnection

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

var conexionsita *sql.DB

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

// func to load .env variables for database
func loadEnvVariables() (DbConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return DbConfig{}, err
	}
	return DbConfig{
		Host:     os.Getenv("host"),
		Port:     os.Getenv("port"),
		User:     os.Getenv("user"),
		Password: os.Getenv("password"),
		DbName:   os.Getenv("dbname"),
	}, nil
}

// Connect_BD conecta con la base de datos y devuelve un error si falla la conexión.
func Connect_BD() error {
	var err error
	dbConfig, err := loadEnvVariables()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)

	conexionsita, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexión exitosa a la base de datos:", conexionsita)
	DataBasePAPA = DbConnection{conexionsita}
	DataBasePAPA.PingOrDie()
	return err
}
