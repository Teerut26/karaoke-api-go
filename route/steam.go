package route

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lrstanley/go-ytdlp"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type SteamRequestBody struct {
	YoutubeURL string `json:"youtube_url"`
}

func SteamHandler(c *fiber.Ctx) error {
	youtubeURL := c.Query("youtube_url")
	height := c.Query("height")
	redirect := c.Query("redirect")

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

	fileName := uuid.New().String()
	fileNameWithUnix := fileName + "_" + strconv.FormatInt(expirsTime.Unix(), 10)
	fileFullName := "video/" + fileNameWithUnix

	ytdlp.MustInstall(context.TODO(), nil)
	dl := ytdlp.New().Format(`bestvideo[height<=` + height + `]+bestaudio[ext=webm][protocol=https]/best`).Output(fileFullName + ".%(ext)s")
	_, err := dl.Run(context.TODO(), youtubeURL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	folderPath := "hls/" + fileNameWithUnix
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0755)
	}
	// ffmpeg -i input.webm -c:v libx264 -c:a aac -hls_time 10 -hls_list_size 0 -f hls output.m3u8
	go func() {
		err = ffmpeg_go.Input(fileFullName+".webm").Output(folderPath+"/"+fileName+".m3u8", ffmpeg_go.KwArgs{"hls_time": 10, "hls_list_size": 0, "c:v": "libx264", "c:a": "aac", "f": "hls"}).OverWriteOutput().Run()
		if err != nil {
			return
		}
	}()

	switch redirect {
	case "video":
		return c.Redirect("/v1/video/" + fileNameWithUnix + ".webm")
	case "hls":
		return c.Redirect("/v1/hls/" + fileNameWithUnix + "/" + fileName + ".m3u8")
	default:
		return c.JSON(fiber.Map{
			"video": `/v1/video/` + fileNameWithUnix + `.webm`,
			"hls":   `/v1/hls/` + fileNameWithUnix + `/` + fileName + `.m3u8`,
		})
	}
}
