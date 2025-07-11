package routes

import (
	"go_fiber/controllers"
	"go_fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	user := router.Group("/users", middleware.JWTProtected(), middleware.AdminOnly())
	user.Get("/", controllers.GetAllUsers)
}
