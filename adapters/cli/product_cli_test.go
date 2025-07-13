package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/viniciusidacruz/hexagonal-archtecture/adapters/cli"
	mock_application "github.com/viniciusidacruz/hexagonal-archtecture/application/mocks"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product 1"
	productPrice := 100.0
	productStatus := "enabled"
	productId := uuid.NewV4().String()

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productService := mock_application.NewMockProductServiceInterface(ctrl)
	productService.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productService.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	productService.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productService.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product created with id %s and name %s with price %f and status %s", productId, productName, productPrice, productStatus)
	result, err := cli.Run(productService, "create", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product enabled with id %s", productId)
	result, err = cli.Run(productService, "enable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product disabled with id %s", productId)
	result, err = cli.Run(productService, "disable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s with id %s", productName, productId)
	result, err = cli.Run(productService, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
