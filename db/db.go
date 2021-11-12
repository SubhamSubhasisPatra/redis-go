package db

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matched record found in the DB")
	Ctx    = context.TODO()
)

func NewDatebase(addrs string) (*Database, error) {
	// mkae the redis coonnection config
	clinet := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := clinet.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: clinet,
	}, nil

}
