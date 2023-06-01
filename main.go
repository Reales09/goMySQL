package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	// fmt.Println(db.ExisteTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	// db.TruncateTable("users")
	// user := models.CreaUser("Adriana", "123456", "Adriana@gmail.com")
	// fmt.Println(user)

	// users := models.ListUsers()
	// fmt.Println(users)

	user := models.GetUser(2)
	fmt.Println(user)

	db.Close()
	// db.Ping()

}
