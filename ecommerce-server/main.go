package main

import (
	"log"

	"github.com/aswinithukku/ecommerce-server/database"
	"github.com/aswinithukku/ecommerce-server/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello world....!")
}

func setupUser(app *fiber.App) {
	app.Get("/api", welcome)

	// User Router
	app.Post("/api/users/create", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
}

func main() {

	database.ConnectDb()
	app := fiber.New()

	setupUser(app)

	log.Fatal(app.Listen(":5050"))
}
