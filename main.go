package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/viniciusidacruz/hexagonal-archtecture/adapters/db"
	"github.com/viniciusidacruz/hexagonal-archtecture/application"
	prismadb "github.com/viniciusidacruz/hexagonal-archtecture/infra/prisma/db"
)

func main() {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL não definida")
	}

	client := prismadb.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		log.Fatalf("erro ao conectar no banco: %v", err)
	}
	defer client.Prisma.Disconnect()

	repo := db.NewProductPrisma(client)
	service := application.NewProductService(repo)

	product, err := service.Create("Product 1", 100)
	if err != nil {
		log.Fatalf("erro ao criar produto: %v", err)
	}

	service.Enable(product)

	found, err := repo.Get(product.GetID())
	if err != nil {
		log.Fatalf("erro ao buscar produto: %v", err)
	}

	log.Printf("Produto criado: %+v", product)
	log.Printf("Produto encontrado: %+v", found)
}
