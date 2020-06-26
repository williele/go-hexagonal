package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve RestFUL API",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("serve")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
