package routes

import (
	"github.com/aswinithukku/ecommerce-server/database"
	"github.com/aswinithukku/ecommerce-server/models"
	"github.com/gofiber/fiber/v2"
)

type UserSerializer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func CreateResponseUser(userModel models.User) UserSerializer {
	return UserSerializer{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName, Email: userModel.Email}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var isEmailExist models.User

	database.Database.Db.Find(&isEmailExist, "email = ?", user.Email)

	if len(isEmailExist.Email) > 0 {
		return c.Status(500).JSON("User already exist!")
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(201).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.Db.Find(&users)

	responseUsers := []UserSerializer{}

	for _, val := range users {
		responseUser := CreateResponseUser(val)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func GetUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(500).JSON("Invalid user id")
	}

	var user models.User

	database.Database.Db.Find(&user, "id = ?", userId)

	if user.ID == 0 {
		return c.Status(500).JSON("User does not found")
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(500).JSON("Invalid user id")
	}

	var user models.User

	database.Database.Db.Find(&user, "id = ?", userId)

	if user.ID == 0 {
		return c.Status(500).JSON("User does not found")
	}

	var body UserSerializer

	if err := c.BodyParser(&body); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	if len(body.FirstName) > 0 {
		user.FirstName = body.FirstName
	}
	if len(body.LastName) > 0 {
		user.LastName = body.LastName
	}

	if len(body.FirstName) <= 0 && len(body.LastName) <= 0 {
		existingUser := CreateResponseUser(user)
		return c.Status(200).JSON(map[string]interface{}{
			"status": "success",
			"user":   existingUser,
		})
	}

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(map[string]interface{}{
		"status": "success",
		"user":   responseUser,
	})

}

func DeleteUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid user id")
	}

	var user models.User

	database.Database.Db.Find(&user, "id = ?", userId)

	if user.ID == 0 {
		return c.Status(400).JSON("User does not found")
	}

	result := database.Database.Db.Delete(&user).Error

	if result != nil {
		return c.Status(404).JSON(result.Error())
	}

	return c.Status(200).SendString("Successfully deleted")

}
