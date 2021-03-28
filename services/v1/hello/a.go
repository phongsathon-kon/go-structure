package hello

import "github.com/gofiber/fiber/v2"

func Greet(c *fiber.Ctx) error {
	return c.SendString("hello world!!")
}