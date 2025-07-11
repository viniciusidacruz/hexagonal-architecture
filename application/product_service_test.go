package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/viniciusidacruz/hexagonal-archtecture/application"
	mock_application "github.com/viniciusidacruz/hexagonal-archtecture/application/mocks"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}

	productResult, err := service.Get("123")
	require.Nil(t, err)
	require.Equal(t, product, productResult)
}
