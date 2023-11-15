package mongo

import (
	"context"
	device "github.com/alserov/device-shop/device-service/pkg/entity"
	"github.com/alserov/device-shop/user-service/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	AddToFavourite(ctx context.Context, req *entity.AddReq) error
	RemoveFromFavourite(ctx context.Context, req *entity.RemoveReq) error
	GetFavourite(ctx context.Context, userUUID string) ([]*device.Device, error)

	AddToCart(ctx context.Context, req *entity.AddReq) error
	RemoveFromCart(ctx context.Context, req *entity.RemoveReq) error
	GetCart(ctx context.Context, userUUID string) ([]*device.Device, error)
}

type repo struct {
	db *mongo.Client
}

func NewRepo(db *mongo.Client) Repository {
	return &repo{
		db: db,
	}
}

func (r repo) AddToFavourite(ctx context.Context, req *entity.AddReq) error {
	coll := r.db.Database("collections").Collection("favourite")

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
	coll := r.db.Database("collections").Collection("favourite")

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
	coll := r.db.Database("collections").Collection("favourite")

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
	coll := r.db.Database("collections").Collection("cart")

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
	coll := r.db.Database("collections").Collection("cart")

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
	coll := r.db.Database("collections").Collection("cart")

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
