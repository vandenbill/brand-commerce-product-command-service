package mongo

import (
	"context"
	"log"

	"github.com/vandenbill/brand-commerce/product-command-service/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepo struct {
	dbClient *mongo.Client
}

func NewProductRepo(dbClient *mongo.Client) domain.ProductRepoMongo {
	return &productRepo{dbClient: dbClient}
}

func (p *productRepo) FindProduct(idPrimitive primitive.ObjectID) (*mongo.SingleResult, error) {
	log.Printf("FindProduct repo invoked")

	coll := p.dbClient.Database("product-service").Collection("product")
	filter := bson.D{{"_id", idPrimitive}}

	result := coll.FindOne(context.Background(), filter)

	return result, nil
}

func (p *productRepo) SaveProduct(data interface{}) (interface{}, error) {
	log.Printf("SaveProduct repo invoked")

	coll := p.dbClient.Database("product-service").Collection("product")
	result, err := coll.InsertOne(context.Background(), data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result.InsertedID, nil
}

func (p *productRepo) EditProduct(idPrimitive primitive.ObjectID, data interface{}) (*mongo.UpdateResult, error) {
	log.Printf("EditProduct repo invoked")

	coll := p.dbClient.Database("product-service").Collection("product")
	filter := bson.D{{"_id", idPrimitive}}

	result, err := coll.ReplaceOne(context.TODO(), filter, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *productRepo) RemoveProduct(idPrimitive primitive.ObjectID) (*mongo.DeleteResult, error) {
	log.Printf("RemoveProduct repo invoked")

	coll := p.dbClient.Database("product-service").Collection("product")
	filter := bson.D{{"_id", idPrimitive}}

	result, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
