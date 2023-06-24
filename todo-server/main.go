package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Println("Hellow world")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Ok working good")
	})

	app.Post("/api/todos/create", func(c *fiber.Ctx) error {
		todo := &Todo{}

		err := c.BodyParser(todo)
		if err != nil {
			return err
		}
		fmt.Println(todo, "todo")

		todo.Id = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)
	})

	app.Get("/api/todos/view", func(c *fiber.Ctx) error {
		todos := todos
		if len(todos) == 0 {
			return c.Status(500).SendString("this has no data in the database")
		}

		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("invalid Id")
		}

		type Res struct {
			Response Todo   `json:"response"`
			Every    []Todo `json:"every"`
			Text     string `json:"text"`
		}
		response := &Res{}

		for i, val := range todos {
			if val.Id == id {
				todos[i].Done = true
				response.Response = todos[i]
				response.Every = todos
				response.Text = "Updated Successfully"
				return c.Status(200).JSON(response)
			}
		}
		return nil
	})

	log.Fatal(app.Listen(":4040"))
}
