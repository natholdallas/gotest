package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	t "github.com/natholdallas/gotest/pkg/tools"
)

type User struct {
	Nickname string         `json:"nickname"`
	Username *string        `json:"username"`
	Password sql.NullString `json:"password"`
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			fmt.Println("Error Handler")
			fmt.Printf("%#v \n", err)
			return err
		},
	})

	app.Get("", func(c *fiber.Ctx) error {
		user := User{}
		c.BodyParser(&user)
		t.PrintJSON(user)
		return c.JSON(user)
	})

	app.Listen(":5000")
}
