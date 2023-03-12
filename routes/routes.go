package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karncomsci/kasetwisai-shop/controller"
)

func Setup(app *fiber.App) {

	app.Get("/", controller.Hello)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
}
