package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			fmt.Println("Error Handler")
			fmt.Printf("%#v \n", err)
			return err
		},
	})

	app.Get("", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK)
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return nil
	})

	app.Listen(":8000")
}
