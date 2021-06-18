package controllers

import (
	"github.com/gleo08/GolangWeek4/database"
	"github.com/gleo08/GolangWeek4/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.JSON(err)
	}
	database.DB.Create(&category)
	return c.JSON(category)
}

func GetAllCategory(c *fiber.Ctx) error {
	var categories []models.Category
	database.DB.Find(&categories)
	return c.JSON(categories)
}

func UpdateCategoryById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	category := models.Category{
		Id: id,
	}
	if err := c.BodyParser(&category); err != nil {
		return c.JSON(err)
	}
	database.DB.Model(&category).Updates(category)
	return c.JSON(category)
}

type UpdateCategoryInput struct {
	Name string `json:"name"`
}

func UpdateInfoCategoryById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var input UpdateCategoryInput
	if err := c.BodyParser(&input); err != nil {
		return c.JSON(err)
	}
	var category models.Category
	if err := database.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": "Cannot find category",
		})
	}
	database.DB.Model(&category).Updates(&models.Category{Name: input.Name})
	return c.JSON(category)
}

func DeleteCategoryById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	database.DB.Exec("DELETE FROM category WHERE id = ?", id)
	return c.JSON(fiber.Map{
		"message": "Delete Successfully",
	})
}
