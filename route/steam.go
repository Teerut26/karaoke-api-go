package route

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lrstanley/go-ytdlp"
)

type SteamRequestBody struct {
	YoutubeURL string `json:"youtube_url"`
}

func SteamHandler(c *fiber.Ctx) error {
	youtubeURL := c.Query("youtube_url")
	height := c.Query("height")

	if height == "" {
		height = "480"
	}

	if youtubeURL == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing youtube_url query parameter")
	}
	if _, err := os.Stat("video"); os.IsNotExist(err) {
		os.Mkdir("video", 0755)
	}

	expirsTime := time.Now().Add(5 * time.Minute)
	fileName := "video/" + uuid.New().String() + "_" + strconv.FormatInt(expirsTime.Unix(), 10)

	ytdlp.MustInstall(context.TODO(), nil)
	dl := ytdlp.New().Format(`bestvideo[height<=` + height + `]+bestaudio[ext=webm][protocol=https]/best`).Output(fileName + ".%(ext)s")
	_, err := dl.Run(context.TODO(), youtubeURL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	videoPath := fileName + ".webm"

	if _, err := os.Stat(videoPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).SendString("Video file not found")
	}

	// defer os.Remove(videoPath)

	c.Set("Content-Type", "video/webm")
	c.Set("Accept-Ranges", "bytes")

	return c.SendFile(videoPath)
}
