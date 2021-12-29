package video

import "fmt"

type Video struct {
	Caption  []interface{} `json:"caption,omitempty"`
	Type     string        `json:"type,omitempty"`
	External struct {
		Url string `json:"url,omitempty"`
	} `json:"external,omitempty"`
}

func ValueOf(s string) Video {
	v := Video{}
	v.Type = "external"
	v.External.Url = fmt.Sprintf("https://www.youtube.com/watch?v=%v", s)

	return v
}
