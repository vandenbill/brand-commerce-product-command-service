package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/vandenbill/brand-commerce-product-command-service/delivery/http"
	"github.com/vandenbill/brand-commerce-product-command-service/repository/mongo"
	"github.com/vandenbill/brand-commerce-product-command-service/usecase"
	"github.com/vandenbill/brand-commerce-product-command-service/util"
)

// TODO implement echo logging
// TODO implement go doc
// TODO implement swagger
// TODO modularize ur code

func main() {
	e := echo.New()

	jaegerHOSTPORT := os.Getenv("JAEGER_HOST_PORT")
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 10,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerHOSTPORT,
		},
	}
	closer, err := cfg.InitGlobalTracer(
		"product-command-service",
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

	rabbitmqURI := os.Getenv("RABBITMQ_URI")
	conn, err := amqp.Dial(rabbitmqURI)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"product-queue", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	util.FailOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbClient, ctx := mongo.NewClient()
	defer func() {
		if err := dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	productRepoMongo := mongo.NewProductRepo(dbClient)
	productUsecase := usecase.NewProductUsecase(productRepoMongo)
	productHandler := http.NewProductHttpDeliver(productUsecase, ctx, q, ch)

	e.POST("/api/product", productHandler.CreateProductHandler)
	e.PUT("/api/product/:id", productHandler.UpdateProductHandler)
	e.DELETE("/api/product/:id", productHandler.DeleteProductHandler)

	log.Fatalln(e.Start("0.0.0.0:8080"))
}
