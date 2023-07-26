package services

import (
  "github.com/atticus64/users-api/src/db"
	userModel "github.com/atticus64/users-api/src/models/user"
	"github.com/gofiber/fiber/v2"
)



func DeleteUser(c *fiber.Ctx) error {
	
	id := c.Params("id")
	
	var user userModel.Model
	result := db.Ctx.First(&user, id)
	
	if result.Error != nil {
		return fiber.NewError(400, "User does not exist!")
	}

	db.Ctx.Delete(&user)	
	
	c.Status(200)
	c.SendString("User created")
	return nil
}


