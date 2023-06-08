package orderRoutes

import (
	"github.com/gofiber/fiber/v2"
	order "project_1/internal/handlers/order"
)

func SetupOrderRoutes(router fiber.Router) {
	orders := router.Group("/order")
	// Create an order
	orders.Post("/", order.CreateOrder)
	// Read all orders
	orders.Get("/", order.GetOrders)
	// // Read one order
	orders.Get("/:order_id", order.GetOrder)
	// // Update one order
	orders.Put("/:order_id", order.UpdateOrder)
	// // Delete one order
	orders.Delete("/:order_id", order.DeleteOrder)
}
