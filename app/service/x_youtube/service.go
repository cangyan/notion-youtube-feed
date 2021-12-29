package x_youtube

import (
	"fmt"
	"log"

	"google.golang.org/api/youtube/v3"
)

type Service interface {
	GetSubscriptionsChannelList() ([]string, error)
	GetChannelPlayListIds(channelId string) ([]string, error)
}
type service struct {
	ApiKey       string
	ClientId     string
	ClientSecret string
	PlayListIds  string
}

func NewService(ak, ci, cs, playListIds string) Service {
	return &service{
		ApiKey:       ak,
		ClientId:     ci,
		ClientSecret: cs,
		PlayListIds:  playListIds,
	}
}

func (s *service) GetSubscriptionsChannelList() ([]string, error) {
	var data []string
	service, err := youtube.New(OAuth2Client(s.ClientId, s.ClientSecret))
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

func (s *service) GetChannelPlayListIds(channelId string) ([]string, error) {
	var data []string
	channelIds, err := s.GetSubscriptionsChannelList()
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	service, err := youtube.New(OAuth2Client(s.ClientId, s.ClientSecret))
	if err != nil {
		return data, err
	}
	for _, channelId := range channelIds {
		call := service.Channels.List([]string{"contentDetails"}).Id(channelId)
		response, err := call.Do()
		if err != nil {
			log.Fatalf("%v", err)
		}

		if len(response.Items) > 0 {
			playList := response.Items[0]
			uploads := playList.ContentDetails.RelatedPlaylists.Uploads

			data = append(data, uploads)

			// call := service.PlaylistItems.List([]string{"snippet"}).PlaylistId(uploads).MaxResults(10)
			// response, err := call.Do()
			// if err != nil {
			// 	log.Fatalf("%v", err)
			// }
			// // b, _ := json.Marshal(response)
			// // fmt.Println(string(b))
			// return response.Items, nil
		}

	}

	return data, nil
}
