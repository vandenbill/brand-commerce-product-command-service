package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/vandenbill/brand-commerce/product-command-service/delivery/http"
	"github.com/vandenbill/brand-commerce/product-command-service/repository/mongo"
	"github.com/vandenbill/brand-commerce/product-command-service/usecase"
)

func main() {
	e := echo.New()

	dbClient, ctx := mongo.NewClient()
	defer func() {
		if err := dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	productRepoMongo := mongo.NewProductRepo(dbClient)
	productUsecase := usecase.NewProductUsecase(productRepoMongo)
	productHandler := http.NewProductHttpDeliver(productUsecase)

	e.POST("/api/product", productHandler.CreateProductHandler)
	e.POST("/api/product", productHandler.CreateProductHandler)
	e.PUT("/api/product/:id", productHandler.UpdateProductHandler)
	e.DELETE("/api/product/:id", productHandler.DeleteProductHandler)

	log.Fatalln(e.Start("0.0.0.0:8080"))
}
