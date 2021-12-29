package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cangyan/notion-youtube-feed/app/service/notion"
	"github.com/cangyan/notion-youtube-feed/utils"
	"github.com/spf13/cobra"
)

var syncVideosToNotionCmd = &cobra.Command{
	Use:   "syncVideosToNotion",
	Short: "同步视频到notion",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("syncVideosToNotion called")

		playListIdsStr := container.Config().YouTubePlayListIds
		playListIds := strings.Split(playListIdsStr, ",")
		service := container.XYouTubeService()
		for _, playListId := range playListIds {
			fmt.Println(playListId)
			list, err := service.GetVideoListByPlayListId(playListId)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if list == nil {
				continue
			}

			// b, _ := json.Marshal(list)
			// fmt.Println(string(b))
			var queryIds []string
			for _, item := range list {
				queryIds = append(queryIds, item.Snippet.ResourceId.VideoId)
			}

			existedIds, err := container.NotionService().FindNotionPageExistedById(queryIds)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, item := range list {
				if utils.StringInArr(item.Snippet.ResourceId.VideoId, existedIds) {
					fmt.Println(item.Snippet.ResourceId.VideoId + "视频已存在")
					continue
				}
				entity := notion.BuildCreateEntity(container.Config().NotionDatabaseId, item)

				b, _ := json.Marshal(entity)
				// fmt.Println(string(b))
				err := container.NotionService().CreatePage(string(b))
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(item.Snippet.ResourceId.VideoId + "视频新建成功")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(syncVideosToNotionCmd)
}
