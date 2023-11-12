package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type Repository interface {
	DeleteWhereExists(ctx context.Context, deviceUUID string) error
}

type repo struct {
	db *mongo.Client
}

func NewRepo(db *mongo.Client) Repository {
	return &repo{
		db: db,
	}
}

const (
	DB_NAME                 = "collections"
	DB_FAVOURITE_COLLECTION = "favourite"
	DB_CART_COLLECTION      = "cart"
)

func (r repo) DeleteWhereExists(ctx context.Context, deviceUUID string) error {
	chErr := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		_, err := r.db.Database(DB_NAME).Collection(DB_FAVOURITE_COLLECTION).DeleteMany(ctx, bson.D{
			{"device.uuid", deviceUUID},
		})
		if err != nil {
			chErr <- err
		}
	}()

	go func() {
		defer wg.Done()
		_, err := r.db.Database(DB_NAME).Collection(DB_CART_COLLECTION).DeleteMany(ctx, bson.D{
			{"device.uuid", deviceUUID},
		})
		if err != nil {
			chErr <- err
		}
	}()
	wg.Wait()
	close(chErr)

	if err := <-chErr; err != nil {
		return err
	}

	return nil
}
