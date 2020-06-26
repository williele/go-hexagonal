package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate PG database",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("migrate")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
