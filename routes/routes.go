package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router){
	AuthRoutes(app)
	NoteRoutes(app)
	ProductRoutes(app)
	UserRoutes(app)
}