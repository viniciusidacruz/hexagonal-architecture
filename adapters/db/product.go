package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/viniciusidacruz/hexagonal-archtecture/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
