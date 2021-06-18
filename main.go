package main

import (
	"github.com/gleo08/GolangWeek4/database"
	"github.com/gleo08/GolangWeek4/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	productRouter := app.Group("/api/product")
	routes.ConfigProductRouter(&productRouter)

	categoryRouter := app.Group("/api/category")
	routes.ConfigCategoryRouter(&categoryRouter)

	userRouter := app.Group("/api/user")
	routes.ConfigUserRouter(&userRouter)

	cartRouter := app.Group("/api/cart")
	routes.ConfigCartRouter(&cartRouter)

	app.Listen(":8080")
}
