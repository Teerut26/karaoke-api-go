package main

import (
	"karaoke-api-go/route"
	"karaoke-api-go/route/hls"
	"karaoke-api-go/route/karaoke"
	"karaoke-api-go/route/ws"
	"karaoke-api-go/services"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cron := gocron.NewScheduler(time.UTC)

	cron.Every("5m").Do(func() {
		services.CleanHLS()
		services.CleanVideo()
	})

	cron.StartAsync()

	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	WebSocketServerSkip := ws.NewWebSocketServer("control/skip")
	go WebSocketServerSkip.HandleMessages()
	app.Get("/ws/control/skip", websocket.New(func(conn *websocket.Conn) {
		ws.ControlSkipHandler(conn, WebSocketServerSkip)
	}))

	WebSocketServerSongEnd := ws.NewWebSocketServer("control/songend")
	go WebSocketServerSongEnd.HandleMessages()
	app.Get("/ws/control/songend", websocket.New(func(conn *websocket.Conn) {
		ws.ControlSongEndHandler(conn, WebSocketServerSongEnd)
	}))

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
	hlsRoute.Get(":id/:playlist", hls.PlaylistHandler)
	hlsRoute.Get(":id/:segment", hls.SegmentHandler)

	karaokeRoute := v1Route.Group("/karaoke")
	karaokeRoute.Get("/queues", karaoke.QueuesHandler)
	karaokeRoute.Delete("/delete", karaoke.DeleteHandler)

	log.Fatal(app.Listen(":3000"))
}
