package cmd

import (
	"demo/pkg/database/pg"
	"demo/pkg/database/pg/migrate"
	"log"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate PG database",
	Run: func(c *cobra.Command, args []string) {
		conn, err := pg.NewConnection()
		if err != nil {
			log.Fatal(err)
		}

		migrate.Migrate(conn, args)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
