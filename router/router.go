package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	orderRoutes "project_1/internal/routes/orderli"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	// Setup the Node Routes
	orderRoutes.SetupOrderRoutes(api)
}
