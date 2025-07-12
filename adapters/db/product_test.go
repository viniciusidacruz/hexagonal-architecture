package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func Setup() {
	Db, _ := sql.Open("sqlite3", ":memory:")
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
