package controllers

import (
	"fmt"
	"go-auth/database"
	"go-auth/modles"

	"github.com/gofiber/fiber/v2"
)

// get all product
func GetProducts(c *fiber.Ctx) error {
	var products []modles.Product
	database.DB.Find(&products)
	return c.JSON(products)
}

// Get product by id
func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	var product modles.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}
	return c.JSON(product)
}

// Create product
func CreateProduct(c *fiber.Ctx) error {
	var product modles.Product

	// 👉 Request Body Print (Raw JSON দেখো)
	body := c.Body()
	fmt.Println("📩 Raw Request Body:", string(body))

	// 👉 JSON Parse Test
	if err := c.BodyParser(&product); err != nil {
		fmt.Println("❌ Body Parsing Error:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}

	// 👉 Parsed Data Print
	fmt.Println("✅ Parsed Product Data:", product)

	// 👉 Database এ Insert করো
	if err := database.DB.Create(&product).Error; err != nil {
		fmt.Println("❌ Database Insert Error:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create product"})
	}

	return c.Status(201).JSON(product)
}

// update product
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product modles.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Save(&product)
	return c.JSON(fiber.Map{
		"message": "Product updated",
		"data":    product,
	})
}

// delete product
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product modles.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}
	database.DB.Delete(&product)
	return c.JSON(fiber.Map{
		"message": "Product deleted",
	})
}
