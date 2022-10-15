package cmd

import (
	"context"
	"log"

	"github.com/Salaton/books-api/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "This command will be used to run the server",
	Long:  `Serve is a command used to run the server`,
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.Router(context.Background())
		err := srv.Run()
		if err != nil {
			log.Printf("an error occurred while running the server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
