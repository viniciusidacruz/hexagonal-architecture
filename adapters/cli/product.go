package cli

import (
	"fmt"

	"github.com/viniciusidacruz/hexagonal-archtecture/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {
	var result string

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)

		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("Product created with id %s and name %s with price %f and status %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)

		if err != nil {
			return "", err
		}

		res, err := service.Enable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product enabled with id %s", res.GetID())
	case "disable":
		product, err := service.Get(productId)

		if err != nil {
			return "", err
		}

		res, err := service.Disable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product disabled with id %s", res.GetID())
	default:
		res, err := service.Get(productId)

		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("Product %s with id %s", res.GetName(), res.GetID())
	}

	return result, nil
}
