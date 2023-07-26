package services

import (
	userDto "github.com/atticus64/users-api/src/controllers/user/dto"
  "github.com/atticus64/users-api/src/db"
	userModel "github.com/atticus64/users-api/src/models/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


func CreateUser(c *fiber.Ctx) error {

	user := new(userDto.CreateUser)
	
  if err := c.BodyParser(user); err != nil {
    return fiber.NewError(500, "Error parsing user")
  }
   
	validate := validator.New()
	err := validate.Struct(user)
	
	if err != nil {
		return fiber.NewError(500, "Not valid user")
	}
	
	var model userModel.Model
	
	model.Age = user.Age
	model.Language = user.Language
	model.Name = user.Name
	
	db.Ctx.Create(&model)
	
	c.Status(201)
	c.SendString("User created")
	return nil


}
