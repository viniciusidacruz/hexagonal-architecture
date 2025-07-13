package cmd

import (
	"github.com/spf13/cobra"
	"github.com/viniciusidacruz/hexagonal-archtecture/adapters/cli"
)

var (
	action       string
	productID    string
	productName  string
	productPrice float64
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Executa ações sobre produtos",
	Run: func(cmd *cobra.Command, args []string) {
		cli.Run(productService, action, productID, productName, productPrice)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "enable ou disable um produto")
	cliCmd.Flags().StringVarP(&productID, "id", "i", "", "ID do produto")
	cliCmd.Flags().StringVarP(&productName, "name", "n", "", "nome do produto")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "preço do produto")
}
