package hls

import "github.com/gofiber/fiber/v2"

func SegmentHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	segment := c.Params("segment")
	filePath := "./hls/" + id + "/" + segment
	return c.SendFile(filePath)
}
