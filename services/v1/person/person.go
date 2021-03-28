package person

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"gosturcture/repository"
)

func Get(c *fiber.Ctx) error {
	repo := repository.NewPersonRepo()
	persons,err := repo.Get()
	if err != nil {
		return err
	}
	b,err := json.Marshal(persons)
	if err != nil {
		return err
	}
	return c.Send(b)
}
