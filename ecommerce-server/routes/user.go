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
