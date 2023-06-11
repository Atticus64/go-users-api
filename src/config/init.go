package config

import (
	"net/http"
	"log"
	Router "github.com/atticus64/users-api/src/controllers"
	User  "github.com/atticus64/users-api/src/controllers/user/services"
	"github.com/atticus64/users-api/src/db"
	"fmt"
)



func Run() {
	db.Connect()
	
	http.HandleFunc("/", Router.Hello)
	http.HandleFunc("/user", User.CreateUser)
	http.HandleFunc("/duser", User.DeleteUser)
	http.HandleFunc("/users", User.GetAllUsers)
	http.HandleFunc("/update-user", User.UpdateUser)

	fmt.Printf("listening on http://localhost:%d\n", 8000)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
