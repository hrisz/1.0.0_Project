package orderRoutes

import (
	"github.com/gofiber/fiber/v2"
	order "project_1/internal/handlers/order"
)

func SetupOrderRoutes(router fiber.Router) {
	user := router.Group("/order")
	// Create a user
	user.Post("/", order.CreateOrder)
	// Read all users
	user.Get("/", order.GetOrders)
	// // Read one user
	user.Get("/:order_id", order.GetOrder)
	// // Update one user
	user.Put("/:order_id", order.UpdateUser)
	// // Delete one user
	user.Delete("/:order_id", order.DeleteOrder)
}
