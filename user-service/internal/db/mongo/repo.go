package mongo

import (
	"context"
	device "github.com/alserov/device-shop/device-service/pkg/entity"
	"github.com/alserov/device-shop/user-service/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type Repository interface {
	AddToFavourite(context.Context, *entity.AddReq) error
	RemoveFromFavourite(context.Context, *entity.RemoveReq) error
	GetFavourite(context.Context, string) ([]*device.Device, error)

	AddToCart(context.Context, *entity.AddReq) error
	RemoveFromCart(context.Context, *entity.RemoveReq) error
	GetCart(context.Context, string) ([]*device.Device, error)

	RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error
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

func (r repo) AddToFavourite(ctx context.Context, req *entity.AddReq) error {
	coll := r.db.Database(DB_NAME).Collection(DB_FAVOURITE_COLLECTION)

	_, err := coll.InsertOne(ctx, bson.M{
		"device":   req.Device,
		"userUUID": req.UserUUID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r repo) RemoveFromFavourite(ctx context.Context, req *entity.RemoveReq) error {
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
	Device *device.Device `bson:"device"`
}

func (r repo) GetFavourite(ctx context.Context, userUUID string) ([]*device.Device, error) {
	coll := r.db.Database(DB_NAME).Collection(DB_FAVOURITE_COLLECTION)

	cur, err := coll.Find(ctx, bson.D{{"userUUID", userUUID}})
	if err != nil {
		return nil, err
	}

	var d []*CollectionRes
	if err = cur.All(ctx, &d); err != nil {
		return nil, err
	}

	var devices []*device.Device

	for _, v := range d {
		devices = append(devices, v.Device)
	}

	return devices, nil
}

func (r repo) AddToCart(ctx context.Context, req *entity.AddReq) error {
	coll := r.db.Database(DB_NAME).Collection(DB_CART_COLLECTION)

	_, err := coll.InsertOne(ctx, bson.M{
		"device":   req.Device,
		"userUUID": req.UserUUID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r repo) RemoveFromCart(ctx context.Context, req *entity.RemoveReq) error {
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

func (r repo) GetCart(ctx context.Context, userUUID string) ([]*device.Device, error) {
	coll := r.db.Database(DB_NAME).Collection(DB_CART_COLLECTION)

	cur, err := coll.Find(ctx, bson.D{{"userUUID", userUUID}})
	if err != nil {
		return nil, err
	}

	var d []*CollectionRes
	if err = cur.All(ctx, &d); err != nil {
		return nil, err
	}

	var devices []*device.Device

	for _, v := range d {
		devices = append(devices, v.Device)
	}

	return devices, nil
}

func (r repo) RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error {
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
