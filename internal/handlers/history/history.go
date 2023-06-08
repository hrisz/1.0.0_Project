package history

import (
	"github.com/gofiber/fiber/v2"
	"project_1/database"
	"project_1/internal/model"
)

//GET ALL Method in history orders

func GetOrders(c *fiber.Ctx) error {
	var orders []tableDB.Order_History
	// Find all orders in database
	result := database.DB.Find(&orders)
	// Check for errors during query execution
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return list of orders
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data History Berhasil Ditampilkan!",
		"data":    orders,
	})
}

//POST Method

func CreateOrder(c *fiber.Ctx) error {
	// Parse request body
	var order tableDB.Order_History
	if err := c.BodyParser(&order); err != nil {
		return err
	}
	// Insert new order into database
	result := database.DB.Create(&order)
	// Check for errors during insertion
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return success message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Order Berhasil Diselesaikan!",
		"data":    order,
	})
}

//GET by ID Method

func GetOrder(c *fiber.Ctx) error {
	// Get order_id parameter from request url
	id := c.Params("order_id")
	// Find order by order_id in database
	var order tableDB.Order_History
	result := database.DB.First(&order, id)
	// Check if order exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Order Tidak Ditermukan!",
		})
	}
	// Check for errors during query
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return order
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    order,
	})
}