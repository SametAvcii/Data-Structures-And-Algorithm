package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
)

type Client struct {
	Client *redis.Client
}

func NewClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

func (c *Client) Get(ctx context.Context, key string) bool {

	err := c.Client.Get(ctx, key).Err()
	if err != nil {
		return false
	}
	return true

}

func (c *Client) Set(ctx context.Context, key string, list interface{}, ex int64) error {
	jsondata, err := json.Marshal(list)
	if err != nil {
		return err
	}
	err = c.Client.Set(ctx, key, jsondata, time.Duration(ex*int64(time.Second))).Err()
	if err != nil {
		return err
	}
	return nil
}
