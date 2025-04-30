package main

import (
	systemLog "log"

	fiberLog "github.com/gofiber/fiber/v2/log"
)

func main() {
	systemLog.Println("System Log")
	fiberLog.Info("Fiber Log")
}
