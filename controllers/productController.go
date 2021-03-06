package controllers

import (
	"fmt"
	_ "strconv"

	"github.com/gleo08/GolangWeek4/database"
	"github.com/gleo08/GolangWeek4/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Preload("Images").Preload("Reviews").Find(&products)
	for i, _ := range products {
		var sum float32 = 0
		var count float32 = 0
		for _, review := range products[i].Reviews {
			sum += review.Rating
			count++
		}
		var averageRating float32
		if sum == 0 || count == 0 {
			averageRating = 0
		} else {
			averageRating = sum / count
		}
		products[i].AverageRating = averageRating
	}
	return c.JSON(products)
}

func GetProductById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var product models.Product
	database.DB.Preload("Images").Preload("Reviews").Where("id = ?", id).First(&product)
	var sum float32 = 0
	var count float32 = 0
	for _, review := range product.Reviews {
		sum += review.Rating
		count++
	}
	var averageRating float32
	if sum == 0 || count == 0 {
		averageRating = 0
	} else {
		averageRating = sum / count
	}
	product.AverageRating = averageRating
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.JSON(err)
	}
	fmt.Println(product)
	database.DB.Create(&product)
	return c.JSON(product)
}

func UpdateProductById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	product := models.Product{
		Id: id,
	}
	if err := c.BodyParser(&product); err != nil {
		return c.JSON(err)
	}
	database.DB.Model(&product).Updates(product)
	return c.JSON(product)
}

type UpdateProductInput struct {
	Price float64 `json:"price"`
}

func UpdateInfoProductById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var input UpdateProductInput
	if err := c.BodyParser(&input); err != nil {
		return c.JSON(err)
	}
	var product models.Product
	if err := database.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": "Cannot find product",
		})
	}
	oldValue := product.Price
	database.DB.Model(&product).Updates(&models.Product{Price: input.Price})
	oldPrice := models.Price{
		ProductId: id,
		Value:     oldValue,
	}
	database.DB.Create(&oldPrice)
	return c.JSON(product)
}

func DeleteProductById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	database.DB.Exec("DELETE FROM product WHERE id = ?", id)
	return c.JSON(fiber.Map{
		"message": "Successfully",
	})
}
