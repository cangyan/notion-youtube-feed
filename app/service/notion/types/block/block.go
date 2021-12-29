package block

import (
	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/video"
	"github.com/google/uuid"
)

type BlockBase struct {
	Object         string `json:"object,omitempty"`
	Id             string `json:"id,omitempty"`
	CreatedTime    string `json:"created_time,omitempty"`
	LastEditedTime string `json:"last_edited_time,omitempty"`
	HasChildren    bool   `json:"has_children,omitempty"`
	Archived       bool   `json:"archived,omitempty"`
	Type           string `json:"type,omitempty"`
}

type VideoBlock struct {
	BlockBase
	Video video.Video `json:"video,omitempty"`
}

func ValueOf(video video.Video) VideoBlock {
	b := VideoBlock{}
	b.Object = "block"
	b.Id = uuid.New().String()
	b.Type = "video"
	b.Video = video

	return b
}
