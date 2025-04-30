package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			fmt.Println("Error Handler")
			fmt.Printf("%#v \n", err)
			return err
		},
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		a := c.Params("id")
		b, err := c.ParamsInt("id", 1)
		log.Info(a)
		log.Info(b)
		if err != nil {
			log.Error(err)
		}
		return nil
	})

	app.Listen(":5000")
}
