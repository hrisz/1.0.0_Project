package order

import (
	"github.com/gofiber/fiber/v2"
	"project_1/database"
	"project_1/internal/model"
)

func GetUsers(c *fiber.Ctx) error {
	var orders []orderTable.Order
	// Find all users in database
	result := database.DB.Find(&orders)
	// Check for errors during query execution
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return list of users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data User Berhasil Ditampilkan!",
		"data":    orders,
	})
}
