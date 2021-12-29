package notion

import (
	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/date"
	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/files"
	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/rich_text"
	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/title"
)

type NotionYouTubePage struct {
	Parent struct {
		DatabaseId string `json:"database_id,omitempty"`
	} `json:"parent,omitempty"`
	Properties NotionYouTubePageProperties `json:"properties,omitempty"`
	Children   []interface{}               `json:"children,omitempty"`
}

type NotionYouTubePageProperties struct {
	Title       title.Title        `json:"标题,omitempty"`
	Id          rich_text.RichText `json:"ID,omitempty"`
	Image       files.Files        `json:"预览图片,omitempty"`
	ChannelName rich_text.RichText `json:"频道名,omitempty"`
	CreatedAt   date.Date          `json:"创建时间,omitempty"`
}

type NotionDatabaseQueryResp struct {
	Object     string      `json:"object,omitempty"`
	Results    []Block     `json:"results,omitempty"`
	NextCursor interface{} `json:"next_cursor,omitempty"`
	HasMore    bool        `json:"has_more,omitempty"`
}

type Block struct {
	Object     string `json:"object,omitempty"`
	Id         string `json:"id,omitempty"`
	Properties struct {
		ID struct {
			RichText []struct {
				Text struct {
					Content string `json:"content,omitempty"`
				} `json:"text,omitempty"`
			} `json:"rich_text,omitempty"`
		} `json:"id,omitempty"`
	} `json:"properties,omitempty"`
}

func (r *NotionDatabaseQueryResp) GetArticleIds() []string {
	data := make([]string, 0)
	if len(r.Results) > 0 {
		for _, item := range r.Results {
			data = append(data, item.Properties.ID.RichText[0].Text.Content)
		}
	}

	return data
}
