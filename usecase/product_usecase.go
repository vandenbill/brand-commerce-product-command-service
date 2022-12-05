package usecase

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/vandenbill/brand-commerce/product-command-service/model/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productUsecase struct {
	productRepoMongo domain.ProductRepoMongo
}

func NewProductUsecase(productRepoMongo domain.ProductRepoMongo) domain.ProductUsecase {
	return &productUsecase{productRepoMongo: productRepoMongo}
}

func (p *productUsecase) CreateProductUsecase(c echo.Context, jaegerCtx context.Context) (interface{}, map[string]interface{}, error) {
	trace, ctx := opentracing.StartSpanFromContext(jaegerCtx, "CreateProductUsecase")
	defer trace.Finish()

	data := map[string]interface{}{}
	if err := c.Bind(&data); err != nil {
		return nil, nil, err
	}

	id, err := p.productRepoMongo.SaveProduct(data, ctx)
	if err != nil {
		return nil, nil, err
	}

	data["id"] = id.(primitive.ObjectID).Hex()
	data["method"] = "create"

	return id, data, nil
}

func (p *productUsecase) UpdateProductUsecase(c echo.Context, jaegerCtx context.Context) (interface{}, interface{}, error) {
	trace, ctx := opentracing.StartSpanFromContext(jaegerCtx, "UpdateProductUsecase")
	defer trace.Finish()

	param := c.Param("id")
	id := strings.Replace(param, "/", "", -1)
	data := map[string]interface{}{}
	if err := c.Bind(&data); err != nil {
		return nil, nil, err
	}

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	result, err := p.productRepoMongo.EditProduct(idPrimitive, data, ctx)
	if err != nil {
		return nil, nil, err
	}
	if result.MatchedCount == 0 {
		return nil, nil, errors.New("error no one data matched")
	}

	data["id"] = id
	data["method"] = "update"

	return result, data, err
}

func (p *productUsecase) DeleteProductUsecase(c echo.Context, jaegerCtx context.Context) (interface{}, string, error) {
	trace, ctx := opentracing.StartSpanFromContext(jaegerCtx, "DeleteProductUsecase")
	defer trace.Finish()

	param := c.Param("id")
	id := strings.Replace(param, "/", "", -1)

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	result, err := p.productRepoMongo.RemoveProduct(idPrimitive, ctx)
	if err != nil {
		return nil, "", err
	}
	if result.DeletedCount == 0 {
		return nil, "", errors.New("cannot delete data")
	}

	return result, id, nil
}
