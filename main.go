package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"go_fiber/config"
	"go_fiber/models"
	"go_fiber/routes"
)

func main(){
	_ = godotenv.Load()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Note{}, &models.Product{})

	app := fiber.New()

	// Inject DB to context

	app.Use(func(c *fiber.Ctx)error{
		c.Locals("db",config.DB)
		return c.Next()
	}) 

	//route group
	api := app.Group("/api")
	routes.SetupRoutes(api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))

	//migrasi db

	

}