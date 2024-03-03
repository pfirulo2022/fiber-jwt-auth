package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pfirulo2022/fiber-jwt-auth/database"
)

func main() {
	database.ConnectDB() // Connect to the database
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",       // Allow all origins
		AllowMethods:     "GET,POST",                    // Allow only GET and POST
		AllowCredentials: true,                          // Check the allowed methods on each
		AllowHeaders:     "Origin, Content-Type,Accept", // Match custom headers
	}))

	log.Fatal(app.Listen(":8000"))
}
