package main

import (
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	db.CreateTable(models.UserSchema)

	db.Close()
	// db.Ping()
}
