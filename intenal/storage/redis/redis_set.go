package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *Redis) Set(ctx context.Context, key string, resp *models.Response) error {

	value, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, value, time.Minute*time.Duration(expirationTime)).Err()
}
