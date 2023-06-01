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
	user := models.CreaUser("Myles", "123456", "Adriana@gmail.com")
	fmt.Println(user)

	// users := models.ListUsers()
	// fmt.Println(users)

	// user := models.GetUser(2)
	// fmt.Println(user)
	// user.Username = "Pedro"
	// user.Password = "P123456"
	// user.Email = "Pedro@gmail.com"
	// user.Save()
	// user.Delete()
	// db.TruncateTable("users")
	fmt.Println(models.ListUsers())

	db.Close()
	// db.Ping()

}
