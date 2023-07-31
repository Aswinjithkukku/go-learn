package routes

import (
	"time"

	"github.com/aswinithukku/ecommerce-server/database"
	"github.com/aswinithukku/ecommerce-server/models"
	"github.com/gofiber/fiber/v2"
)

type OrderSerializer struct {
	ID        uint             `json:"id"`
	User      UserSerializer   `json:"user"`
	Product   ProductSeializer `json:"product"`
	CreatedAt time.Time        `json:"createdAt"`
}

func CreateResponseOrder(orderModel models.Order, user UserSerializer, product ProductSeializer) OrderSerializer {
	return OrderSerializer{ID: orderModel.ID, User: user, Product: product, CreatedAt: orderModel.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	database.Database.Db.Find(&user, "id = ?", order.UserReferer)

	if user.ID == 0 {
		return c.Status(500).JSON("Cannot find User")
	}

	var product models.Product

	database.Database.Db.Find(&product, "id = ?", order.ProductReferer)

	if product.ID == 0 {
		return c.Status(500).JSON("Cannot find product")
	}

	database.Database.Db.Create(&order)
	if order.ID == 0 {
		return c.Status(404).JSON("Order does not found")
	}
	userResponse := CreateResponseUser(user)
	productResponse := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, userResponse, productResponse)

	return c.Status(201).JSON(responseOrder)
}
