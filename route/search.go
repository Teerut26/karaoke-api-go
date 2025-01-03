package route

import "github.com/gofiber/fiber/v2"

func SearchHandler(c *fiber.Ctx) error {
	return c.SendString("Search")

}
