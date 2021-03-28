package rest

import (
	"github.com/gofiber/fiber/v2"
	"gosturcture/services/middleware"
	"gosturcture/services/v1/hello"
	"gosturcture/services/v1/person"
)

func Router(app fiber.Router) {
	v1 := app.Group("/service/v1",middleware.Interceptor)
	v1.Get("/greet",hello.Greet)
	v1.Get("/person",person.Get)
}