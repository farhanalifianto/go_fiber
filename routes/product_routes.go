package routes

import (
	"go_fiber/controllers"
	"go_fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router fiber.Router) {
	product := router.Group("/products")

	product.Get("/", controllers.GetAllProducts)
	product.Post("/", middleware.JWTProtected(), controllers.CreateProduct)
	product.Put("/:id", middleware.JWTProtected(),controllers.UpdateProduct)
	product.Delete("/:id", middleware.JWTProtected(),controllers.DeleteProduct)
}
