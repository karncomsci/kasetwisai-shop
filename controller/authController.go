package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karncomsci/kasetwisai-shop/db"
	"github.com/karncomsci/kasetwisai-shop/models"
	"golang.org/x/crypto/bcrypt"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello world!")
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)

	user := models.User{
		Username: data["username"],
		Password: password,
		FullName: data["fullname"],
		Avatar:   data["avatar"],
	}

	db.DB.Create(&user)

	return c.JSON(user)

}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	println("password: ", data["password"])
	println("avatar: ", data["avatar"])

	var user models.User

	db.DB.Where("avatar = ?", data["avatar"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	var test = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	println("ttttt :", test)
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	return c.JSON(user)
}
