package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	dbInfra "github.com/viniciusidacruz/hexagonal-archtecture/adapters/db"
	"github.com/viniciusidacruz/hexagonal-archtecture/application"
	prismadb "github.com/viniciusidacruz/hexagonal-archtecture/infra/prisma/db"
)

var (
	prismaClient   *prismadb.PrismaClient
	productService *application.ProductService
)

var rootCmd = &cobra.Command{
	Use:   "hexagonal-architecture",
	Short: "CLI da aplicação hexagonal-architecture",
}

func Execute() {
	prismaClient = prismadb.NewClient()
	if err := prismaClient.Prisma.Connect(); err != nil {
		log.Fatalf("erro ao conectar no banco: %v", err)
	}
	defer prismaClient.Prisma.Disconnect()

	repo := dbInfra.NewProductPrisma(prismaClient)
	productService = application.NewProductService(repo)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "help message for toggle")
}
