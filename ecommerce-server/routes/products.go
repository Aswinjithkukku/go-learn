package routes

import (
	"github.com/aswinithukku/ecommerce-server/database"
	"github.com/aswinithukku/ecommerce-server/models"
	"github.com/gofiber/fiber/v2"
)

type ProductSeializer struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serialNumber"`
}

func CreateResponseProduct(productModel models.Product) ProductSeializer {
	return ProductSeializer{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.Database.Db.Find(&products)

	if len(products) == 0 {
		return c.Status(500).JSON("cannot find the products")
	}

	return c.Status(200).JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product

	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Invalid productId")
	}

	database.Database.Db.Find(&product, "id = ?", productId)

	if product.ID == 0 {
		return c.Status(500).JSON("Cannot find product")
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	var body struct {
		Name         string
		SerialNumber string
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON("Cannot parse body")
	}

	var product models.Product

	productId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid productId")
	}

	database.Database.Db.Find(&product, "id = ?", productId)

	if product.ID == 0 {
		return c.Status(500).JSON("Cannot find product")
	}

	product.Name = body.Name
	product.SerialNumber = body.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)

}

func DeleteProduct(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Find(&product, "id == ?", userId)

	if product.ID == 0 {
		return c.Status(500).JSON("cannot find the product")
	}

	result := database.Database.Db.Delete(&product).Error

	if result != nil {
		return c.Status(404).JSON(result.Error())
	}

	return c.Status(200).JSON("Product deleted successfully")
}
