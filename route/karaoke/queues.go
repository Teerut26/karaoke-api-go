package karaoke

import (
	"context"
	"encoding/json"
	"karaoke-api-go/config"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

type response struct {
	ID         string `json:"id"`
	FireBaseID string `json:"fireBaseId"`
	Channel    string `json:"channel"`
	CreatedAt  int    `json:"createdAt"`
	Thumbnail  string `json:"thumbnail"`
	Title      string `json:"title"`
	URL        string `json:"url"`
}

func QueuesHandler(c *fiber.Ctx) error {
	context := context.Background()
	client := config.Firestore()
	defer client.Close()

	iter := client.Collection("songs").OrderBy("createdAt", firestore.Asc).Documents(context)
	datas := []response{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
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

		datas = append(datas, data)
	}

	return c.JSON(datas)
}
