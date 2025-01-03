package main

import (
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
	v1Route.Get("/sources", route.SourcesHandler)
	v1Route.Get("/search", route.SearchHandler)
	v1Route.Get("/video/:file", route.VideoHandler)

	hlsRoute := v1Route.Group("/hls")
	route.HLSHandler(hlsRoute)

	app.Listen(":3000")
}
