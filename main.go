package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pfirulo2022/fiber-jwt-auth/database"
	"github.com/pfirulo2022/fiber-jwt-auth/router"
)

func main() {
	database.ConnectDB() // Connect to the database
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost",            // Allow all origins
		AllowMethods:     "GET,POST",                    // Allow only GET and POST
		AllowCredentials: true,                          // Check the allowed methods on each
		AllowHeaders:     "Origin, Content-Type,Accept", // Match custom headers
	}))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
