package route

import "github.com/gofiber/fiber/v2"

func PlaylistHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	playlist := c.Params("playlist")
	filePath := "./hls/" + id + "/" + playlist
	return c.SendFile(filePath)
}
