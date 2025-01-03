package route

import "github.com/gofiber/fiber/v2"

func HLSHandler(c fiber.Router) error {
	c.Get(":id/:playlist", playlistHandler)
	c.Get(":id/:segment", segmentHandler)
	return nil
}

func playlistHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	playlist := c.Params("playlist")
	filePath := "./hls/" + id + "/" + playlist
	return c.SendFile(filePath)
}

func segmentHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	segment := c.Params("segment")
	filePath := "./hls/" + id + "/" + segment

	return c.SendFile(filePath)
}
