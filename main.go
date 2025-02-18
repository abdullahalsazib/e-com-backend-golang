package main

import (
	"go-auth/database"
	"go-auth/routes"
	"log"

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
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))

	// run the server
	routes.Setup(app)
	routes.SetupProductRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
