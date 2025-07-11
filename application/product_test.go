package application_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/viniciusidacruz/hexagonal-archtecture/application"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Price = 100
	product.Status = application.ProductStatusDisabled

	err := product.Enable()

	require.Nil(t, err)
	require.Equal(t, product.Status, application.ProductStatusEnabled)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}
