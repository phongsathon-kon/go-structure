package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Interceptor(c *fiber.Ctx) error {
	fmt.Println("im interceptor")
	return c.Next()
}