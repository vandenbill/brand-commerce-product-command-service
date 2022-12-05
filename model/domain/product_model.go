package domain

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepoMongo interface {
	FindProduct(id primitive.ObjectID) (*mongo.SingleResult, error)
	SaveProduct(data interface{}) (interface{}, error)
	EditProduct(id primitive.ObjectID, data interface{}) (*mongo.UpdateResult, error)
	RemoveProduct(id primitive.ObjectID) (*mongo.DeleteResult, error)
}

type ProductUsecase interface {
	CreateProductUsecase(c echo.Context) (interface{}, map[string]interface{}, error)
	UpdateProductUsecase(c echo.Context) (interface{}, interface{}, error)
	DeleteProductUsecase(c echo.Context) (interface{}, string, error)
}

type ProductHttpDeliver interface {
	CreateProductHandler(c echo.Context) error
	UpdateProductHandler(c echo.Context) error
	DeleteProductHandler(c echo.Context) error
}
