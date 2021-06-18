package routes

import (
	"github.com/gleo08/GolangWeek4/controllers"
	"github.com/gofiber/fiber/v2"
)

func ConfigProductRouter(router *fiber.Router) {
	(*router).Get("/", controllers.GetAllProducts)
	(*router).Post("/", controllers.CreateProduct)
	(*router).Delete("/:id", controllers.DeleteProductById)
	(*router).Put("/:id", controllers.UpdateProductById)
	(*router).Patch("/:id", controllers.UpdateInfoProductById)
}

func ConfigCategoryRouter(router *fiber.Router) {
	(*router).Post("/", controllers.CreateCategory)
	(*router).Get("/", controllers.GetAllCategory)
	(*router).Put("/:id", controllers.UpdateCategoryById)
	(*router).Patch("/:id", controllers.UpdateInfoCategoryById)
	(*router).Delete("/:id", controllers.DeleteCategoryById)
}

func ConfigUserRouter(router *fiber.Router) {
	(*router).Post("/", controllers.CreateUser)
	(*router).Get("/", controllers.GetAllUsers)
	(*router).Get("/:id", controllers.GetUserById)
	(*router).Put("/:id", controllers.UpdateUserById)
	(*router).Patch("/:id", controllers.UpdateInfoUserById)
	(*router).Delete("/:id", controllers.DeleteUserById)
}

func ConfigCartRouter(router *fiber.Router) {
	(*router).Get("/", controllers.GetAllCarts)
	(*router).Post("/", controllers.CreateCart)
}
