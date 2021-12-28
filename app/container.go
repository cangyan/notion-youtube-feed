package app

import (
	"os"

	"github.com/cangyan/notion-youtube-feed/app/service/x_youtube"
)

type Container struct {
	config          *Config
	xYouTubeService x_youtube.Service
}

func (c *Container) Config() *Config {
	if c.config == nil {
		notionToken := os.Getenv("NOTION_TOKEN")
		if notionToken == "" {
			panic("NOTION_TOKEN not config")
		}

		notionDatabaseId := os.Getenv("NOTION_DATABASE_ID")
		if notionDatabaseId == "" {
			panic("NOTION_DATABASE_ID not config")
		}

		youtubeApiKey := os.Getenv("YOUTUBE_API_KEY")
		if youtubeApiKey == "" {
			panic("YOUTUBE_API_KEY not config")
		}

		youtubeClientId := os.Getenv("YOUTUBE_CLIENT_ID")
		if youtubeClientId == "" {
			panic("YOUTUBE_CLIENT_ID not config")
		}

		youtubeClientSecret := os.Getenv("YOUTUBE_CLIENT_SECRET")
		if youtubeClientSecret == "" {
			panic("YOUTUBE_CLIENT_SECRET not config")
		}

		c.config = &Config{
			NotionToken:         notionToken,
			NotionDatabaseId:    notionDatabaseId,
			YouTubeApiKey:       youtubeApiKey,
			YouTubeClientId:     youtubeClientId,
			YouTubeClientSecret: youtubeClientSecret,
		}
	}

	return c.config
}

func (c *Container) XYouTubeService() x_youtube.Service {
	if c.xYouTubeService == nil {
		c.xYouTubeService = x_youtube.NewService(c.Config().YouTubeApiKey, c.Config().YouTubeClientId, c.Config().YouTubeClientSecret)
	}

	return c.xYouTubeService
}
