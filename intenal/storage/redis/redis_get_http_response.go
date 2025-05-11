package redis

import (
	"context"
	"encoding/json"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *Redis) Get(ctx context.Context, key string) (*models.Response, error) {
	resp := &models.Response{}

	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(value), resp); err != nil {
		return nil, err
	}

	return resp, nil
}
