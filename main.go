package main

import (
	"fmt"
	"karaoke-api-go/route"
	"karaoke-api-go/services"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cron := gocron.NewScheduler(time.UTC)

	cron.Every("5m").Do(func() {
		services.CleanHLS()
		services.CleanVideo()
	})

	cron.StartAsync()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	v1Route := app.Group("/v1")
	v1Route.Get("/steam", route.SteamHandler)
	v1Route.Get("/video/:file", route.VideoHandler)

	hlsRote := v1Route.Group("/hls")
	hlsRote.Get(":id/:playlist", func(c *fiber.Ctx) error {
		id := c.Params("id")
		playlist := c.Params("playlist")

		fmt.Println("id", id)
		fmt.Println("playlist", playlist)
		filePath := "./hls/" + id + "/" + playlist
		return c.SendFile(filePath)
	})
	hlsRote.Get(":id/:segment", func(c *fiber.Ctx) error {
		id := c.Params("id")
		segment := c.Params("segment")
		filePath := "./hls/" + id + "/" + segment

		return c.SendFile(filePath)
	})

	app.Listen(":3000")
}
