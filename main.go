package main

import (
	"github.com/gofiber/fiber/v2"
	"rest-go/config"
	"rest-go/routes"
)

func main() {
	app := fiber.New()

	//run database
	config.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":6000")
}
