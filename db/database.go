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

//Verifica si una tabla existe o no

func ExisteTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()

}

// Crea una tabla

func CreateTable(schema string, name string) {

	if !ExisteTable(name) {
		_, err := db.Exec(schema)

		if err != nil {

			fmt.Println(err)
		}
	}

}

// Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}

	return result, err

}

// Polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	if err != nil {
		fmt.Println(err)
	}

	return rows, err

}
