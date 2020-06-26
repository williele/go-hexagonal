package cmd

import (
	"demo/pkg/database/pg"
	"demo/pkg/database/pg/repository"
	"demo/pkg/server/rest"
	"demo/pkg/services/categories"
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

		categoryRepo := repository.NewCategoryRepository(conn)
		categoryService := categories.NewService(categoryRepo)

		server := rest.NewHTTPRest(":3000", productService, categoryService)
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
