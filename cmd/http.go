package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"template-2025-feb/internal/config"
)

// httpCmd represents the command to start the Fiber HTTP server
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start the Fiber HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig(".env")
		config.InitLogger()
		config.InitDB()

		config.Logger.Info("Starting HTTP server...")
		startHTTPServer()
	},
}

func startHTTPServer() {
	app := fiber.New()

	port := viper.GetString("HTTP_PORT")
	if port == "" {
		port = "9090"
	}

	config.Logger.Fatal(app.Listen(":" + port))
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
