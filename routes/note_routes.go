package routes

import (
	"go_fiber/controllers"
	"go_fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func NoteRoutes(router fiber.Router) {
	note := router.Group("/notes", middleware.JWTProtected())

	note.Get("/", controllers.GetNotes)
	note.Post("/", controllers.CreateNote)
	note.Put("/:id", controllers.UpdateNote)
	note.Delete("/:id", controllers.DeleteNote)
	note.Post("/:id/favorite", controllers.ToggleFavorite)
	note.Get("/favorites", controllers.GetFavoriteNotes)	
}
