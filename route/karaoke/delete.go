package karaoke

import (
	"context"
	"karaoke-api-go/config"

	"github.com/gofiber/fiber/v2"
)

type requestSchema struct {
	ID string `json:"id"`
}

type responseSchema struct {
	ID string `json:"id"`
}

func DeleteHandler(c *fiber.Ctx) error {
	context := context.Background()

	body := new(requestSchema)
	if err := c.BodyParser(body); err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	client := config.Firestore()
	defer client.Close()

	_, err := client.Collection("songs").Doc(body.ID).Delete(context)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(responseSchema{
		ID: body.ID,
	})
}
