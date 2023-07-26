package services

import (
	userDto "github.com/atticus64/users-api/src/controllers/user/dto"
  "github.com/atticus64/users-api/src/db"
	userModel "github.com/atticus64/users-api/src/models/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userBody := new(userDto.UpdateUser)
	
  if err := c.BodyParser(userBody); err != nil {
    return fiber.NewError(500, "Error parsing user")
  }
  validate := validator.New()
	err := validate.Struct(userBody)
	
	if err != nil {
		return fiber.NewError(500, "Not valid user")
	}

  var user userModel.Model
	result := db.Ctx.First(&user, id)
	
	if result.Error != nil {
		return fiber.NewError(400, "User does not exist!")
	}	
  
  if (userBody.Age != 0){
  	user.Age = userBody.Age	
  } 
  
  if (userBody.Name != "") {
  	user.Name = userBody.Name	
  }
  
  if (userBody.Language != "") {
  	user.Language = userBody.Language	
  }
 
	db.Ctx.Save(&user)
		
	c.SendString("User Updated")
	return nil
}
