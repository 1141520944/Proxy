package redis

import (
	"context"
	"time"
)

func (r *Redis) Put_string(c context.Context, key, value string, expire time.Duration) error {
	return r.DB.Set(c, key, value, expire).Err()
}
func (r *Redis) Get_string(c context.Context, key string) (string, error) {
	return r.DB.Get(c, key).Result()
}
