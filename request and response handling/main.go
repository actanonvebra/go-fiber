package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/users/:id", func(c *fiber.Ctx) error {

		userID := c.Params("id")

		return c.SendString("User ID: " + userID)
	})

	// app.Get("/sayHi/:name", func(c *fiber.Ctx) error {

	// 	userID := c.Params("name")

	// 	return c.SendString("Hello " + userID)
	// })

	app.Get("/sayHi/json/", func(c *fiber.Ctx) error {

		response := fiber.Map{
			"message": "Hello",
		}

		return c.JSON(response)
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		err := someFunctionThatMayFail()
		if err != nil {
			panic(err)
		}
		return c.SendString("No error occured")
	})

	app.Use(func(c *fiber.Ctx) error {
		if err := recover(); err != nil {
			return c.Status(500).SendString("Internal Server Error")
		}
		return c.Next()
	})

	app.Listen(":3000")

}
