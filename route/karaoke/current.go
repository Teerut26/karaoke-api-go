package karaoke

import (
	"context"
	"encoding/json"
	"karaoke-api-go/config"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
)

func CurrentHandler(c *fiber.Ctx) error {
	context := context.Background()
	client := config.Firestore()
	defer client.Close()

	iter := client.Collection("songs").OrderBy("createdAt", firestore.Asc).Documents(context)
	doc, err := iter.Next()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No song in queue",
		})
	}

	data := response{}
	docData, err := json.Marshal(doc.Data())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	err = json.Unmarshal(docData, &data)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	data.FireBaseID = doc.Ref.ID

	return c.JSON(data)
}
