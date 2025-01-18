package route

import (
	"context"
	"encoding/json"
	"karaoke-api-go/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/lrstanley/go-ytdlp"
)

func SourcesHandler(c *fiber.Ctx) error {
	youtubeURL := c.Query("youtube_url")
	height := c.Query("height")

	if height == "" {
		height = "720"
	}

	if youtubeURL == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing youtube_url query parameter")
	}

	ytdlp.MustInstall(context.TODO(), nil)
	dl := ytdlp.New().Format(`bestvideo[height<=` + height + `]+bestaudio[ext=webm][protocol=https]/best`).DumpJSON()
	// dl := ytdlp.New().Format(`bestvideo[height<=` + height + `]+bestaudio/best`).ExtractorArgs("youtube:player_client=ios").DumpJSON()
	res, err := dl.Run(context.TODO(), youtubeURL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	result := new(dto.TYDLPResponse)
	json.Unmarshal([]byte(res.Stdout), &result)
	return c.Status(200).JSON(result.RequestedFormats)
}
