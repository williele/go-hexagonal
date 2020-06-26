package cmd

import (
	"demo/pkg/database/pg"
	"demo/pkg/database/pg/repository"
	"demo/pkg/server/rest"
	"demo/pkg/services/products"
	"log"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve RestFUL API",
	Run: func(c *cobra.Command, args []string) {
		conn, err := pg.NewConnection()
		if err != nil {
			log.Fatal(err)
		}

		productRepo := repository.NewProductRepository(conn)
		productService := products.NewService(productRepo)

		server := rest.NewHTTPRest(":3000", productService)
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
