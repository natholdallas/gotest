package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natholdallas/natools4go/ptr"
)

func Err(value any, statis ...int) *fiber.Error {
	msg := "unknown"
	code := fiber.StatusBadRequest
	if str, ok := value.(string); ok {
		msg = str
	} else if err, ok := value.(error); ok {
		msg = err.Error()
	}
	return &fiber.Error{Code: code, Message: msg}
}

func main() {
	err := Err("hello world")
	ptr.JSON(err)
}
