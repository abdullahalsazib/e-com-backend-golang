package main

import (
	"fmt"
	"go-auth/database"
	"go-auth/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	database.Connection()
}
func main() {

	app := fiber.New()
	// router
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, https://your-frontend.vercel.app", // Allow local and deployed frontend
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// run the server
	routes.Setup(app)
	routes.SetupProductRoutes(app)
	// Get PORT from Render Environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not provided
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
