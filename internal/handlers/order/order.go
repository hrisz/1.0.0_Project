package order

import (
	"github.com/gofiber/fiber/v2"
	"project_1/database"
	"project_1/internal/model"
)

//GET ALL Method

func GetUsers(c *fiber.Ctx) error {
	var orders []orderTable.Order_Table
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

//POST Method

func CreateUser(c *fiber.Ctx) error {
	// Parse request body
	var user orderTable.Order_Table
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	// Insert new user into database
	result := database.DB.Create(&user)
	// Check for errors during insertion
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return success message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Berhasil Ditambahkan!",
		"data":    user,
	})
}

//GET by ID Method

func GetUser(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	id := c.Params("id_user")
	// Find user by id_user in database
	var user orderTable.Order_Table
	result := database.DB.First(&user, id)
	// Check if user exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Tidak Ditermukan!",
		})
	}
	// Check for errors during query
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return user
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    user,
	})
}

// PUT/UPDATE Method

func UpdateUser(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	id := c.Params("id_user")
	// Find user by id_user in database
	var user orderTable.Order_Table
	result := database.DB.First(&user, id)
	// Check if user exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Tidak Ditemukan",
		})
	}
	// Parse request body
	var newUser orderTable.Order_Table
	if err := c.BodyParser(&newUser); err != nil {
		return err
	}
	// Update user in database
	result = database.DB.Model(&user).Updates(newUser)
	// Check for errors during update
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User Berhasil Diperbarui!",
		"data":    user,
	})
}

//DELETE Method

func DeleteUser(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	id := c.Params("id_user")
	// Find user by id_user in database
	var user orderTable.Order_Table
	result := database.DB.First(&user, id)
	// Check if user exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Tidak Ditemukan",
		})
	}
	// Delete user from database
	result = database.DB.Delete(&user)
	// Check for errors during deletion
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User Berhasil Dihapus!",
		"data":    user,
	})
}
