package usecase

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/vandenbill/brand-commerce/product-command-service/model/domain"
	"log"
	"strings"
)

type productUsecase struct {
	productRepoMongo domain.ProductRepoMongo
}

func NewProductUsecase(productRepoMongo domain.ProductRepoMongo) domain.ProductUsecase {
	return &productUsecase{productRepoMongo: productRepoMongo}
}

func (p *productUsecase) CreateProductUsecase(c echo.Context) (interface{}, error) {
	log.Printf("CreateProduct service invoked")

	data := map[string]interface{}{}
	if err := c.Bind(&data); err != nil {
		return nil, err
	}

	id, err := p.productRepoMongo.SaveProduct(data)
	if err != nil {
		return nil, err
	}

	return id, err
}

func (p *productUsecase) UpdateProductUsecase(c echo.Context) (interface{}, error) {
	log.Printf("UpdateProduct service invoked")

	param := c.Param("id")
	id := strings.Replace(param, "/", "", -1)
	data := map[string]interface{}{}
	if err := c.Bind(&data); err != nil {
		return nil, err
	}

	result, err := p.productRepoMongo.EditProduct(id, data)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("Error no one data matched")
	}

	return result, err
}

func (p *productUsecase) DeleteProductUsecase(c echo.Context) (interface{}, error) {
	log.Printf("DeleteProduct service invoked")

	param := c.Param("id")
	id := strings.Replace(param, "/", "", -1)

	result, err := p.productRepoMongo.RemoveProduct(id)
	if err != nil {
		return nil, err
	}
	if result.DeletedCount == 0 {
		return nil, errors.New("Cannot delete data.")
	}

	return result, nil
}
