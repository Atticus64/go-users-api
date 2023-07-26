package services

import (
	"github.com/atticus64/users-api/src/db"
	userModel "github.com/atticus64/users-api/src/models/user"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	
	users := []*userModel.Model{} 
	
	db.Ctx.Find(&users)
	
	c.JSON(users)
	c.Status(200)
	return nil
}