package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"github.com/viniciusidacruz/hexagonal-archtecture/application"
)

var Db *sql.DB

func Setup() {
	var err error
	Db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	)`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products values ("abc", "Product 1", 100, "enabled")`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	Setup()

	defer Db.Close()

	productDb := NewProductDb(Db)

	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 100.00, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	Setup()

	defer Db.Close()

	productDb := NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product 2"
	product.Price = 200

	productCreated, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.GetID(), productCreated.GetID())
	require.Equal(t, product.GetName(), productCreated.GetName())
	require.Equal(t, product.GetPrice(), productCreated.GetPrice())
	require.Equal(t, product.GetStatus(), productCreated.GetStatus())

	product.Status = "enabled"

	productUpdated, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.GetID(), productUpdated.GetID())
	require.Equal(t, product.GetName(), productUpdated.GetName())
	require.Equal(t, product.GetPrice(), productUpdated.GetPrice())
	require.Equal(t, product.GetStatus(), productUpdated.GetStatus())
}
