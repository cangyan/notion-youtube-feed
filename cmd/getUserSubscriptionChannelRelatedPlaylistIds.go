package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var getUserSubscriptionChannelRelatedPlaylistIdsCmd = &cobra.Command{
	Use:   "getPlayListIds",
	Short: "获取订阅频道的播放列表集合",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getPlayListIds called")
		service := container.XYouTubeService()
		list, err := service.GetChannelPlayListIds()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(strings.Join(list, ","))
	},
}

func init() {
	rootCmd.AddCommand(getUserSubscriptionChannelRelatedPlaylistIdsCmd)
}
