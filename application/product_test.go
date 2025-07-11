package application_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
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

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Price = 0
	product.Status = application.ProductStatusEnabled

	err := product.Disable()
	require.Nil(t, err)
	require.Equal(t, product.Status, application.ProductStatusDisabled)

	product.Price = 100
	err = product.Disable()
	require.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Price = 100
	product.Status = application.ProductStatusEnabled

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "status must be enabled or disabled", err.Error())

	product.Status = ""
	_, _ = product.IsValid()
	require.Equal(t, product.Status, application.ProductStatusDisabled)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "price must be greater than zero", err.Error())

	product.Price = 0
	_, err = product.IsValid()
	require.Nil(t, err)
}
