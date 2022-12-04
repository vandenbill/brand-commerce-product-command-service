package domain

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepoMongo interface {
	SaveProduct(data interface{}) (interface{}, error)
	EditProduct(id string, data interface{}) (*mongo.UpdateResult, error)
	RemoveProduct(id string) (*mongo.DeleteResult, error)
}

type ProductUsecase interface {
	CreateProductUsecase(c echo.Context) (interface{}, error)
	UpdateProductUsecase(c echo.Context) (interface{}, error)
	DeleteProductUsecase(c echo.Context) (interface{}, error)
}

type ProductHttpDeliver interface {
	CreateProductHandler(c echo.Context) error
	UpdateProductHandler(c echo.Context) error
	DeleteProductHandler(c echo.Context) error
}
