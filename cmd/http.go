package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"log"
)

// httpCmd represents the command to start the Fiber HTTP server
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start the Fiber HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		startHTTPServer()
	},
}

func startHTTPServer() {
	app := fiber.New()

	port := "9090"

	log.Fatal(app.Listen(":" + port))
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
