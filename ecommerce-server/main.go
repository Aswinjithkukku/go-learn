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
	app.Get("/api/users/:id", routes.GetUser)
	app.Patch("/api/users/:id/update", routes.UpdateUser)
	app.Delete("/api/users/:id/delete", routes.DeleteUser)

	// Product Routes
	app.Post("/api/products/create", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Patch("/api/products/:id/update", routes.UpdateProduct)
	app.Delete("/api/products/:id/delete", routes.DeleteProduct)

	// Order Routes
	app.Post("/api/orders/create", routes.CreateOrder)
}

func main() {

	database.ConnectDb()
	app := fiber.New()

	setupUser(app)

	log.Fatal(app.Listen(":5050"))
}
