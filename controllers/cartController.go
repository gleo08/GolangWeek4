package controllers

import (
	"github.com/gleo08/GolangWeek4/database"
	"github.com/gleo08/GolangWeek4/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllCarts(c *fiber.Ctx) error {
	var carts []models.Cart

	database.DB.Find(&carts)
	return c.JSON(carts)
}

func CreateCart(c *fiber.Ctx) error {
	var cart models.Cart
	if err := c.BodyParser(&cart); err != nil {
		return c.JSON(err)
	}
	database.DB.Create(&cart)
	return c.JSON(cart)
}
