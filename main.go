package main

import (
	"BookingPlatform/configs"
	"BookingPlatform/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	routes.UserRoute(app)
	err := app.Listen(":6000")
	if err != nil {
		return
	}
}
