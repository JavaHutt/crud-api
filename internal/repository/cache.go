package repository

import (
	"bytes"
	"context"
	"encoding/gob"
	"strconv"
	"time"

	"github.com/JavaHutt/crud-api/internal/model"

	"github.com/go-redis/redis/v9"
)

const defaultExpiration = 25 * time.Second

type (
	Option func(c *cache)
	cache  struct {
		client *redis.Client
		exp    time.Duration
	}
)

func NewCache(client *redis.Client, options ...Option) cache {
	c := cache{
		client: client,
		exp:    defaultExpiration,
	}

	for _, opt := range options {
		opt(&c)
	}

	return c
}

// Get gets query to the cache
func (c cache) Get(ctx context.Context, id string) (*model.SlowestQuery, error) {
	cmd := c.client.Get(ctx, id)

	cmdb, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}

	b := bytes.NewReader(cmdb)
	var query model.SlowestQuery
	if err = gob.NewDecoder(b).Decode(&query); err != nil {
		return nil, err
	}

	return &query, nil
}

// Set saves query to the cache
func (c cache) Set(ctx context.Context, query *model.SlowestQuery) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(query); err != nil {
		return err
	}

	return c.client.Set(ctx, strconv.Itoa(int(query.ID)), b.Bytes(), c.exp).Err()
}

func WithExpiration(exp time.Duration) Option {
	return func(c *cache) {
		c.exp = exp
	}
}
