package db

import (
	"context"
	"database/sql"

	"github.com/viniciusidacruz/hexagonal-archtecture/application"
	prismadb "github.com/viniciusidacruz/hexagonal-archtecture/infra/prisma/db"
)

type ProductPrisma struct {
	client *prismadb.PrismaClient
}

func NewProductPrisma(client *prismadb.PrismaClient) *ProductPrisma {
	return &ProductPrisma{client: client}
}

func (p *ProductPrisma) Get(id string) (application.ProductInterface, error) {
	ctx := context.Background()

	pr, err := p.client.Product.FindUnique(
		prismadb.Product.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	if pr == nil {
		return nil, sql.ErrNoRows
	}

	product := &application.Product{
		ID:     pr.ID,
		Name:   pr.Name,
		Price:  pr.Price,
		Status: pr.Status,
	}

	return product, nil
}

func (p *ProductPrisma) Save(prod application.ProductInterface) (application.ProductInterface, error) {
	ctx := context.Background()

	pr, err := p.client.Product.
		FindUnique(
			prismadb.Product.ID.Equals(prod.GetID()),
		).
		Update(
			prismadb.Product.Name.Set(prod.GetName()),
			prismadb.Product.Price.Set(prod.GetPrice()),
			prismadb.Product.Status.Set(prod.GetStatus()),
		).
		Exec(ctx)

	if err == prismadb.ErrNotFound {
		pr, err = p.client.Product.CreateOne(
			prismadb.Product.Name.Set(prod.GetName()),
			prismadb.Product.Price.Set(prod.GetPrice()),
			prismadb.Product.Status.Set(prod.GetStatus()),
		).Exec(ctx)
	}

	if err != nil {
		return nil, err
	}

	return &application.Product{
		ID:     pr.ID,
		Name:   pr.Name,
		Price:  pr.Price,
		Status: pr.Status,
	}, nil
}
