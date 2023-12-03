package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MustConnect(ctx context.Context, uri string) *mongo.Client {
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic("failed to connect to db: " + err.Error())
	}

	if err = cl.Ping(ctx, readpref.Primary()); err != nil {
		panic("failed to ping db: " + err.Error())
	}

	return cl
}
