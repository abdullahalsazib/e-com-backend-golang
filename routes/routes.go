package routes

import (
	"go-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Authentication Routes
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/update", controllers.UpdateProfile)
	app.Get("/api/profile", controllers.GetUserProfile)
	app.Get("/api/test", controllers.TestApi)
	// add jack

	// Serve uploaded images
	app.Static("/uploads", "./uploads")

	// Setup Product Routes
	SetupProductRoutes(app)
}

func SetupProductRoutes(app *fiber.App) {
	product := app.Group("/api/products") // Plural naming convention

	product.Get("/", controllers.GetProducts)         // Get all products
	product.Get("/:id", controllers.GetProductById)   // Get product by ID
	product.Post("/", controllers.CreateProduct)      // Create product
	product.Put("/:id", controllers.UpdateProduct)    // Update product
	product.Delete("/:id", controllers.DeleteProduct) // Delete product (Fixed)
}
