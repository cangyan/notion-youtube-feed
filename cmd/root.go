package cmd

import (
	"fmt"
	"os"

	"github.com/cangyan/notion-youtube-feed/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "job",
	Short: "youtube data api tool",
	Long:  `基于cobra定制(cli模式)youtube操作工具`,
}
var container = app.Container{}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
