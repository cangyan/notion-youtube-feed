package main

import (
	"fmt"
	"log"

	"github.com/cangyan/notion-youtube-feed/app"
)

func main() {
	container := &app.Container{}

	youtubeService := container.XYouTubeService()

	list, _ := youtubeService.GetSubscriptionsChannelList()
	fmt.Println(list)
	for _, cid := range list[:1] {
		v, err := youtubeService.GetChannelVideoList(cid)
		if err != nil {
			log.Fatalln(err)
			continue
		}

		for _, item := range v {
			fmt.Println(item.Snippet.Title)
		}
	}
}
