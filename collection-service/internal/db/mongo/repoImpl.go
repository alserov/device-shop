package mongo

import (
	"context"
	"errors"
	"github.com/alserov/device-shop/collection-service/internal/db"
	"github.com/alserov/device-shop/collection-service/internal/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"sync"
)

type repo struct {
	log *slog.Logger
	db  *mongo.Client
}

func NewCollectionsRepo(db *mongo.Client, log *slog.Logger) db.CollectionsRepo {
	return &repo{
		log: log,
		db:  db,
	}
}

const (
	DB_NAME                 = "collections"
	DB_FAVOURITE_COLLECTION = "favourite"
	DB_CART_COLLECTION      = "cart"

	internalError = "internal error"
	notFound      = "nothing found"
)

func (r *repo) AddToFavourite(ctx context.Context, userUUID string, device models.Device) error {
	op := "repo.AddToFavourite"

	coll := r.db.Database(DB_NAME).Collection(DB_FAVOURITE_COLLECTION)

	_, err := coll.InsertOne(ctx, bson.M{
		"device":   device,
		"userUUID": userUUID,
	})
	if errors.Is(mongo.ErrNoDocuments, err) {
		return status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to insert device", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}

	return nil
}

func (r *repo) RemoveFromFavourite(ctx context.Context, userUUID string, deviceUUID string) error {
	op := "repo.RemoveFromFavourite"

	coll := r.db.Database(DB_NAME).Collection(DB_FAVOURITE_COLLECTION)

	_, err := coll.DeleteOne(ctx, bson.D{
		{"userUUID", userUUID},
		{"device.uuid", deviceUUID},
	})
	if errors.Is(mongo.ErrNoDocuments, err) {
		return status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to delete from collection", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}

	return nil
}

func (r *repo) GetFavourite(ctx context.Context, userUUID string) ([]*models.Device, error) {
	op := "repo.GetFavourite"

	coll := r.db.Database(DB_NAME).Collection(DB_FAVOURITE_COLLECTION)

	cur, err := coll.Find(ctx, bson.D{{"userUUID", userUUID}})
	if errors.Is(mongo.ErrNoDocuments, err) {
		return nil, status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to find devices", slog.String("error", err.Error()), slog.String("op", op))
		return nil, status.Error(codes.Internal, internalError)
	}

	var d []*models.DeviceFromCollection
	if err = cur.All(ctx, &d); err != nil {
		r.log.Error("failed to scan devices", slog.String("error", err.Error()), slog.String("op", op))
		return nil, err
	}

	var devices []*models.Device
	for _, v := range d {
		devices = append(devices, &v.Device)
	}

	return devices, nil
}

func (r *repo) AddToCart(ctx context.Context, userUUID string, device models.Device) error {
	op := "repo.AddToCart"

	coll := r.db.Database(DB_NAME).Collection(DB_CART_COLLECTION)

	_, err := coll.InsertOne(ctx, bson.M{
		"device":   device,
		"userUUID": userUUID,
	})
	if errors.Is(mongo.ErrNoDocuments, err) {
		return status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed insert device", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}
	return nil
}

func (r *repo) RemoveFromCart(ctx context.Context, userUUID string, deviceUUID string) error {
	op := "repo.RemoveFromCart"

	coll := r.db.Database(DB_NAME).Collection(DB_CART_COLLECTION)

	_, err := coll.DeleteOne(ctx, bson.D{
		{"userUUID", userUUID},
		{"device.uuid", deviceUUID},
	})
	if errors.Is(mongo.ErrNoDocuments, err) {
		return status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to delete from collection", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}

	return nil
}

func (r *repo) GetCart(ctx context.Context, userUUID string) ([]*models.Device, error) {
	op := "repo.GetCart"

	coll := r.db.Database(DB_NAME).Collection(DB_CART_COLLECTION)

	cur, err := coll.Find(ctx, bson.D{{"userUUID", userUUID}})
	if errors.Is(mongo.ErrNoDocuments, err) {
		return nil, status.Error(codes.NotFound, notFound)
	}
	if err != nil {
		r.log.Error("failed to find devices", slog.String("error", err.Error()), slog.String("op", op))
		return nil, status.Error(codes.Internal, internalError)
	}

	var d []*models.DeviceFromCollection
	if err = cur.All(ctx, &d); err != nil {
		r.log.Error("failed to scan devices", slog.String("error", err.Error()), slog.String("op", op))
		return nil, status.Error(codes.Internal, internalError)
	}

	var devices []*models.Device
	for _, v := range d {
		devices = append(devices, &v.Device)
	}

	return devices, nil
}

func (r *repo) RemoveDeviceFromCollections(ctx context.Context, deviceUUID string) error {
	op := "repo.RemoveDeviceFromCollections"

	chErr := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		_, err := r.db.Database(DB_NAME).Collection(DB_FAVOURITE_COLLECTION).DeleteMany(ctx, bson.D{
			{"device.uuid", deviceUUID},
		})
		if errors.Is(mongo.ErrNoDocuments, err) {
			chErr <- status.Error(codes.NotFound, notFound)
		}
		if err != nil {
			chErr <- status.Error(codes.Internal, internalError)
		}
	}()

	go func() {
		defer wg.Done()
		_, err := r.db.Database(DB_NAME).Collection(DB_CART_COLLECTION).DeleteMany(ctx, bson.D{
			{"device.uuid", deviceUUID},
		})
		if errors.Is(mongo.ErrNoDocuments, err) {
			chErr <- status.Error(codes.NotFound, notFound)
		}
		if err != nil {
			chErr <- status.Error(codes.Internal, internalError)
		}
	}()

	go func() {
		wg.Wait()
		close(chErr)
	}()

	if err := <-chErr; err != nil {
		r.log.Error("failed to remove device from collection", slog.String("error", err.Error()), slog.String("op", op))
		return status.Error(codes.Internal, internalError)
	}

	return nil
}
