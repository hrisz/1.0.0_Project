package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"project_1/database"
	"project_1/router"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	//Connect to Database
	database.ConnectDB()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	//Router Setup
	router.SetupRoutes(app)

	//Listen on PORT 3000
	app.Listen(":3000")
}
