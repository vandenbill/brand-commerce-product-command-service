package mongo

import (
	"context"
	"github.com/vandenbill/brand-commerce/product-command-service/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type productRepo struct {
	dbClient *mongo.Client
}

func NewProductRepo(dbClient *mongo.Client) domain.ProductRepoMongo {
	return &productRepo{dbClient: dbClient}
}

func (p *productRepo) SaveProduct(data interface{}) (interface{}, error) {
	// TODO setup log
	log.Printf("SaveProduct repo invoked")

	coll := p.dbClient.Database("product-service").Collection("product")
	result, err := coll.InsertOne(context.Background(), data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result.InsertedID, nil
}

func (p *productRepo) EditProduct(id string, data interface{}) (*mongo.UpdateResult, error) {
	log.Printf("EditProduct repo invoked")

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	coll := p.dbClient.Database("product-service").Collection("product")
	filter := bson.D{{"_id", idPrimitive}}

	result, err := coll.ReplaceOne(context.TODO(), filter, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *productRepo) RemoveProduct(id string) (*mongo.DeleteResult, error) {
	log.Printf("RemoveProduct repo invoked")

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	coll := p.dbClient.Database("product-service").Collection("product")
	filter := bson.D{{"_id", idPrimitive}}

	result, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
