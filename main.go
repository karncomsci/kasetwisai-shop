package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karncomsci/kasetwisai-shop/db"
	"github.com/karncomsci/kasetwisai-shop/routes"
)

func main() {
	db.DBConnection()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8082")
}
