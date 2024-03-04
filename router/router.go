package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pfirulo2022/fiber-jwt-auth/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/signup", handlers.SignUp)
	app.Post("/signin", handlers.SignIn)
	app.Get("/logout", handlers.Logout)
	app.Get("/user", handlers.User)
}
