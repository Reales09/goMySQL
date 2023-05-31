package main

import (
	"fmt"
	"gomysql/db"
)

func main() {
	db.Connect()
	fmt.Println(db.ExisteTable("users"))
	//db.CreateTable(models.UserSchema, "users")

	db.Close()
	// db.Ping()
}
