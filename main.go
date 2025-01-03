package main

import (
	"karaoke-api-go/route"
	"karaoke-api-go/services"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cron := gocron.NewScheduler(time.UTC)

	cron.Every("5m").Do(services.CleanVideo)

	cron.StartAsync()

	app := fiber.New()

	v1Route := app.Group("/v1")
	v1Route.Get("/steam", route.SteamHandler)
	v1Route.Get("/video/:file", route.VideoHandler)

	app.Listen(":3000")
}
