package notion

import (
	"encoding/json"
	"fmt"

	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/block"
	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/date"
	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/files"
	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/rich_text"
	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/title"
	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/video"
	"google.golang.org/api/youtube/v3"
)

func BuildCreateEntity(databaseId string, item *youtube.PlaylistItem) NotionYouTubePage {
	b, _ := json.Marshal(item)
	fmt.Println(string(b))
	n := NotionYouTubePage{}
	n.Parent.DatabaseId = databaseId
	{
		n.Properties = NotionYouTubePageProperties{
			Title:       title.ValueOf(item.Snippet.Title),
			Id:          rich_text.ValueOf(item.Snippet.ResourceId.VideoId),
			Image:       files.ValueOf(item.Snippet.Thumbnails.Default.Url),
			ChannelName: rich_text.ValueOf(item.Snippet.ChannelTitle),
			CreatedAt:   date.ValueOf(item.Snippet.PublishedAt),
		}
	}

	{
		n.Children = make([]interface{}, 0)
		video := video.ValueOf(item.Snippet.ResourceId.VideoId)
		block := block.ValueOf(video)
		n.Children = append(n.Children, block)
	}
	return n
}
