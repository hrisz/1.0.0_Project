package historyRoutes

import (
	"github.com/gofiber/fiber/v2"
	history "project_1/internal/handlers/history"
)

func SetupHistoryRoutes(router fiber.Router) {
	orders := router.Group("/history")
	// Create an order
	orders.Post("/", history.CreateOrder)
	// Read all orders
	orders.Get("/", history.GetOrders)
	// // Read one order
	orders.Get("/:order_id", history.GetOrder)
}
