package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// GET request handler
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// POST request handler
	app.Post("/data", func(c *fiber.Ctx) error {
		// Read the JSON data from the request body
		type Data struct {
			Message string `json:"message"`
		}
		data := new(Data)
		if err := c.BodyParser(data); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON data")
		}

		// Process the received data
		message := data.Message
		response := fmt.Sprintf("Received message: %s", message)

		return c.SendString(response)
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
