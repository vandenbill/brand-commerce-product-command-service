package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vandenbill/brand-commerce/product-command-service/model/domain"
	"github.com/vandenbill/brand-commerce/product-command-service/model/web"
	"net/http"
)

type productHttpDeliver struct {
	productUsecase domain.ProductUsecase
}

func NewProductHttpDeliver(productUsecase domain.ProductUsecase) domain.ProductHttpDeliver {
	return &productHttpDeliver{productUsecase: productUsecase}
}

func (p *productHttpDeliver) CreateProductHandler(c echo.Context) error {
	id, err := p.productUsecase.CreateProductUsecase(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.BuildErrorResponse("Error could not save data", err.Error()))
		return nil
	}
	c.JSON(http.StatusCreated, web.BuildResponse("Succes create product", id))
	return nil
}

func (p *productHttpDeliver) UpdateProductHandler(c echo.Context) error {
	result, err := p.productUsecase.UpdateProductUsecase(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.BuildErrorResponse("Error could not update data.", err.Error()))
		return nil
	}
	c.JSON(http.StatusCreated, web.BuildResponse("Succes update product", result))
	return nil
}

func (p *productHttpDeliver) DeleteProductHandler(c echo.Context) error {
	result, err := p.productUsecase.DeleteProductUsecase(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.BuildErrorResponse("Error could not remove data.", err.Error()))
		return nil
	}
	c.JSON(http.StatusCreated, web.BuildResponse("Succes remove product.", result))
	return nil
}
