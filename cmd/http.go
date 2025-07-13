package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusidacruz/hexagonal-archtecture/adapters/web/server"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "sobe API HTTP",
	Run: func(cmd *cobra.Command, args []string) {
		if productService == nil {
			fmt.Println("productService não inicializado")
			return
		}

		srv := server.NewWebServer(productService)
		fmt.Println("Server is running on port 8080")
		srv.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
