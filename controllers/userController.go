package controllers

import (
	"github.com/gleo08/GolangWeek4/database"
	"github.com/gleo08/GolangWeek4/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	var user models.User
	id, _ := c.ParamsInt("id")
	err := database.DB.Where("id = ?", id).First(&user)
	if err != nil {
		c.JSON(fiber.Map{
			"message": "Cannot find user",
		})
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		c.JSON(fiber.Map{
			"message": "Cannot parse data",
		})
	}

	database.DB.Create(&user)
	return c.JSON(user)
}

func UpdateUserById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user := models.User{
		Id: id,
	}
	if err := c.BodyParser(&user); err != nil {
		c.JSON(fiber.Map{
			"message": "Cannot parse data",
		})
	}
	database.DB.Model(&user).Updates(user)
	return c.JSON(user)
}

type UpdateUserInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
}

func UpdateInfoUserById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var input UpdateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.JSON(err)
	}
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": "Cannot find user",
		})
	}
	database.DB.Model(&user).Updates(&models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		UserName:  input.UserName,
	})
	return c.JSON(user)
}

func DeleteUserById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	database.DB.Exec("DELETE FROM user WHERE id = ?", id)
	return c.JSON(fiber.Map{
		"message": "Delete successfully",
	})
}
