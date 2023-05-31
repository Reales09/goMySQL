package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username: password@tcp(localhost:3306)/database

const url = "root:123456@tcp(localhost:3306)/goweb_db"

// Guarda la conexión
var db *sql.DB

// Realiza la conexión
func Connect() {
	conection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Conexion exitosa")
	db = conection
}

// Cerrar la conexión
func Close() {
	db.Close()
}

//Verificar la conexión

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
