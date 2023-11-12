package cache

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

type Set struct {
	Key string
	Val interface{}
}

type Repository interface {
	SetValue(ctx context.Context, cred *Set) error
	GetValue(ctx context.Context, key string) (interface{}, error)
}

type repo struct {
	redis *redis.Client
}

func NewRepo(r *redis.Client) Repository {
	return &repo{
		redis: r,
	}
}

func (r *repo) SetValue(ctx context.Context, cred *Set) error {
	var err error
	b := new(bytes.Buffer)
	if err = json.NewEncoder(b).Encode(cred.Val); err != nil {
		return err
	}

	return r.redis.Set(cred.Key, b.Bytes(), time.Hour).Err()
}

func (r *repo) GetValue(ctx context.Context, key string) (interface{}, error) {
	cmd := r.redis.Get(key)

	cmdb, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}

	var res interface{}

	if err = json.Unmarshal(cmdb, &res); err != nil {
		return nil, err
	}

	return res, nil
}
