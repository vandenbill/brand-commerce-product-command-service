package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient() (*mongo.Client, context.Context) {
	// const uri = "mongodb://root:root@localhost:27017/?maxPoolSize=20&w=majority"
	uri := os.Getenv("MONGO_URI")

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	return client, ctx
}
