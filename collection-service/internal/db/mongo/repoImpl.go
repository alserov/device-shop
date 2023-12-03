package mongo

import (
	"context"
	"github.com/alserov/device-shop/collection-service/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type repo struct {
	db *mongo.Client
}

func NewCollectionsRepo(db *mongo.Client) db.CollectionsRepo {
	return &repo{
		db: db,
	}
}

const (
	DB_NAME                 = "collections"
	DB_FAVOURITE_COLLECTION = "favourite"
	DB_CART_COLLECTION      = "cart"
)

func (r *repo) AddToFavourite(ctx context.Context, userUUID string, device db.Device) error {
	coll := r.db.Database(DB_NAME).Collection(DB_FAVOURITE_COLLECTION)

	_, err := coll.InsertOne(ctx, bson.M{
		"device":   device,
		"userUUID": userUUID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) RemoveFromFavourite(ctx context.Context, req *db.ChangeCollectionReq) error {
	coll := r.db.Database(DB_NAME).Collection(DB_FAVOURITE_COLLECTION)

	_, err := coll.DeleteOne(ctx, bson.D{
		{"userUUID", req.UserUUID},
		{"device.uuid", req.DeviceUUID},
	})
	if err != nil {
		return err
	}

	return nil
}

type CollectionRes struct {
	Device db.Device `bson:"device"`
}

func (r *repo) GetFavourite(ctx context.Context, userUUID string) ([]*db.Device, error) {
	coll := r.db.Database(DB_NAME).Collection(DB_FAVOURITE_COLLECTION)

	cur, err := coll.Find(ctx, bson.D{{"userUUID", userUUID}})
	if err != nil {
		return nil, err
	}

	var d []*CollectionRes
	if err = cur.All(ctx, &d); err != nil {
		return nil, err
	}

	var devices []*db.Device
	for _, v := range d {
		devices = append(devices, &v.Device)
	}

	return devices, nil
}

func (r *repo) AddToCart(ctx context.Context, userUUID string, device db.Device) error {
	coll := r.db.Database(DB_NAME).Collection(DB_CART_COLLECTION)

	_, err := coll.InsertOne(ctx, bson.M{
		"device":   device,
		"userUUID": userUUID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) RemoveFromCart(ctx context.Context, req *db.ChangeCollectionReq) error {
	coll := r.db.Database(DB_NAME).Collection(DB_CART_COLLECTION)

	_, err := coll.DeleteOne(ctx, bson.D{
		{"userUUID", req.UserUUID},
		{"device.uuid", req.DeviceUUID},
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetCart(ctx context.Context, userUUID string) ([]*db.Device, error) {
	coll := r.db.Database(DB_NAME).Collection(DB_CART_COLLECTION)

	cur, err := coll.Find(ctx, bson.D{{"userUUID", userUUID}})
	if err != nil {
		return nil, err
	}

	var d []*CollectionRes
	if err = cur.All(ctx, &d); err != nil {
		return nil, err
	}

	var devices []*db.Device

	for _, v := range d {
		devices = append(devices, &v.Device)
	}

	return devices, nil
}

func (r *repo) RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error {
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
