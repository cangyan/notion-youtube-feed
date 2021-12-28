package x_youtube

import (
	"errors"
	"log"

	"google.golang.org/api/youtube/v3"
)

type Service interface {
	GetSubscriptionsChannelList() ([]string, error)
	GetChannelVideoList(channelId string) ([]*youtube.PlaylistItem, error)
}
type service struct {
	ApiKey       string
	ClientId     string
	ClientSecret string
}

func NewService(ak, ci, cs string) Service {
	return &service{
		ApiKey:       ak,
		ClientId:     ci,
		ClientSecret: cs,
	}
}

func (s *service) GetSubscriptionsChannelList() ([]string, error) {
	var data []string
	service, err := youtube.New(GoogleDefaultClient())
	if err != nil {
		log.Fatalf("%v", err)
	}
	call := service.Subscriptions.List([]string{"id", "snippet", "contentDetails"}).Mine(true).MaxResults(10)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("%v", err)
	}

	// b, _ := json.Marshal(response)
	// fmt.Println(string(b))

	for _, item := range response.Items {
		data = append(data, item.Snippet.ResourceId.ChannelId)
		// fmt.Println(item.Snippet.Title)
	}

	return data, nil
}

func (s *service) GetChannelVideoList(channelId string) ([]*youtube.PlaylistItem, error) {
	var data []*youtube.PlaylistItem
	service, err := youtube.New(GoogleDefaultClient())
	if err != nil {
		log.Fatalf("%v", err)
	}
	call := service.Channels.List([]string{"contentDetails"}).Id(channelId)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("%v", err)
	}

	if len(response.Items) > 0 {
		playList := response.Items[0]
		uploads := playList.ContentDetails.RelatedPlaylists.Uploads
		call := service.PlaylistItems.List([]string{"snippet"}).PlaylistId(uploads).MaxResults(10)
		response, err := call.Do()
		if err != nil {
			log.Fatalf("%v", err)
		}
		// b, _ := json.Marshal(response)
		// fmt.Println(string(b))
		return response.Items, nil
	}

	return data, errors.New("videos not found")
}
