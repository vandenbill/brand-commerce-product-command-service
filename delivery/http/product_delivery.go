package http

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rabbitmq/amqp091-go"
	"github.com/vandenbill/brand-commerce-product-command-service/model/domain"
	"github.com/vandenbill/brand-commerce-product-command-service/model/web"
	"github.com/vandenbill/brand-commerce-product-command-service/util"
)

type productHttpDeliver struct {
	productUsecase domain.ProductUsecase
	ctx            context.Context
	q              amqp091.Queue
	ch             *amqp091.Channel
}

func NewProductHttpDeliver(productUsecase domain.ProductUsecase, ctx context.Context, q amqp091.Queue, ch *amqp091.Channel) domain.ProductHttpDeliver {
	return &productHttpDeliver{productUsecase: productUsecase, ctx: ctx, q: q, ch: ch}
}

func (p *productHttpDeliver) CreateProductHandler(c echo.Context) error {
	trace, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "CreateProductHandler")
	defer trace.Finish()

	id, data, err := p.productUsecase.CreateProductUsecase(c, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.BuildErrorResponse("Error could not save data", err.Error()))
		return nil
	}

	c.JSON(http.StatusCreated, web.BuildResponse("Succes create product", id))

	err = p.ch.PublishWithContext(p.ctx,
		"",       // exchange
		p.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})
	util.FailOnError(err, "Failed to publish a message")
	log.Println("Send data")

	return nil
}

func (p *productHttpDeliver) UpdateProductHandler(c echo.Context) error {
	trace, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "UpdateProductHandler")
	defer trace.Finish()

	result, data, err := p.productUsecase.UpdateProductUsecase(c, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.BuildErrorResponse("Error could not update data.", err.Error()))
		return nil
	}

	c.JSON(http.StatusCreated, web.BuildResponse("Succes update product", result))

	err = p.ch.PublishWithContext(p.ctx,
		"",       // exchange
		p.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})
	util.FailOnError(err, "Failed to publish a message")
	log.Println("Send data")

	return nil
}

func (p *productHttpDeliver) DeleteProductHandler(c echo.Context) error {
	trace, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "UpdateProductHandler")
	defer trace.Finish()

	result, id, err := p.productUsecase.DeleteProductUsecase(c, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.BuildErrorResponse("Error could not remove data.", err.Error()))
		return nil
	}
	c.JSON(http.StatusCreated, web.BuildResponse("Succes remove product.", result))

	err = p.ch.PublishWithContext(p.ctx,
		"",       // exchange
		p.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(id),
		})
	util.FailOnError(err, "Failed to publish a message")
	log.Println("Send data")

	return nil
}
