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

const defaultExpiration = 25

type cache struct {
	client *redis.Client
	exp    time.Duration
}

func NewCache(client *redis.Client) cache {
	return cache{
		client: client,
		exp:    defaultExpiration,
	}
}

func (c cache) Get(ctx context.Context, id string) (*model.Advertise, error) {
	cmd := c.client.Get(ctx, id)

	cmdb, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}

	b := bytes.NewReader(cmdb)
	var ad model.Advertise
	if err = gob.NewDecoder(b).Decode(&ad); err != nil {
		return nil, err
	}

	return &ad, nil
}

func (c cache) Set(ctx context.Context, ad model.Advertise) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(ad); err != nil {
		return err
	}

	return c.client.Set(ctx, strconv.Itoa(int(ad.ID)), b.Bytes(), c.exp*time.Second).Err()
}
