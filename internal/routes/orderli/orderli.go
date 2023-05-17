package orderRoutes

import (
	"github.com/gofiber/fiber/v2"
	order "project_1/internal/handlers/order"
)

func SetupOrderRoutes(router fiber.Router) {
	user := router.Group("/order")
	// Create a user
	user.Post("/", order.CreateUser)
	// Read all users
	user.Get("/", order.GetUsers)
	// // Read one user
	user.Get("/:id_user", order.GetUser)
	// // Update one user
	user.Put("/:id_user", order.UpdateUser)
	// // Delete one user
	user.Delete("/:id_user", order.DeleteUser)
}
