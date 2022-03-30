package routes

import (
	"rest-go/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controller.CreateUser)
	app.Get("/user/:userId", controller.GetAUser)
	app.Put("/user/:userId", controller.EditAUser)
	app.Delete("/user/:userId", controller.DeleteAUser)
}
