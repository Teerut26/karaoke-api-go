package route

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func VideoHandler(c *fiber.Ctx) error {
	file := c.Params("file")
	if file == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing file query parameter")
	}

	videoPath := "video/" + file + ".webm"

	if _, err := os.Stat(videoPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).SendString("Video file not found")
	}

	c.Set("Content-Type", "video/webm")
	c.Set("Accept-Ranges", "bytes")

	return c.SendFile(videoPath)
}
