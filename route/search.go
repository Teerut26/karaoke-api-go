package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raitonoberu/ytsearch"
)

type ResponseSearch struct {
	Videos []Video `json:"videos"`
}

type Video struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
	Channel   string `json:"channel"`
}

func SearchHandler(c *fiber.Ctx) error {
	keywords := c.Query("keywords")
	search := ytsearch.VideoSearch(keywords)
	result, err := search.Next()
	if err != nil {
		panic(err)
	}

	var videos []Video
	for _, video := range result.Videos {
		videos = append(videos, Video{
			ID:        video.ID,
			Title:     video.Title,
			Url:       video.URL,
			Thumbnail: video.RichThumbnail.URL,
			Channel:   video.Channel.Title,
		})
	}

	return c.Status(200).JSON(ResponseSearch{
		Videos: videos,
	})
}
