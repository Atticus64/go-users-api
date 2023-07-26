package config

import (
	"fmt"
	"log"
	"net/http"
	Router "github.com/atticus64/users-api/src/controllers"
	User "github.com/atticus64/users-api/src/controllers/user/services"
	"github.com/atticus64/users-api/src/db"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()
	
	db.Connect()
    
  app.Get("/", Router.Hello)
	app.Post("/user", User.CreateUser)
	app.Get("/users", User.GetUsers)
	app.Delete("/user/:id", User.DeleteUser)
	app.Put("/user/:id", User.UpdateUser)
	// http.HandleFunc("/update-user", User.UpdateUser)

	app.Listen(":8000")
	fmt.Printf("listening on http://localhost:%d\n", 8000)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
