package cmd

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A command to start http server",
	Long:  `A command t start http server that use go fiber library`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("created by aprianfirlanda@gmail.com")
		httpServer := figure.NewFigure("Http Server", "big", true)
		fmt.Println(httpServer.String())
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
