package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func Connect(ctx context.Context, uri string) (*mongo.Client, error) {
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err = cl.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	log.Println("mongo connected")
	return cl, nil
}
