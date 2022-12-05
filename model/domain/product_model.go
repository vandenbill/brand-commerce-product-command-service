package domain

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepoMongo interface {
	FindProduct(id primitive.ObjectID, jaegerCtx context.Context) (*mongo.SingleResult, error)
	SaveProduct(data interface{}, jaegerCtx context.Context) (interface{}, error)
	EditProduct(id primitive.ObjectID, data interface{}, jaegerCtx context.Context) (*mongo.UpdateResult, error)
	RemoveProduct(id primitive.ObjectID, jaegerCtx context.Context) (*mongo.DeleteResult, error)
}

type ProductUsecase interface {
	CreateProductUsecase(c echo.Context, jaegerCtx context.Context) (interface{}, []byte, error)
	UpdateProductUsecase(c echo.Context, jaegerCtx context.Context) (interface{}, []byte, error)
	DeleteProductUsecase(c echo.Context, jaegerCtx context.Context) (interface{}, []byte, error)
}

type ProductHttpDeliver interface {
	CreateProductHandler(c echo.Context) error
	UpdateProductHandler(c echo.Context) error
	DeleteProductHandler(c echo.Context) error
}
