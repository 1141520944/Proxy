package repo

import (
	"context"
	"time"
)

//redis

type Cache_String interface {
	Put_string(c context.Context, key, value string, expire time.Duration) error
	Get_string(c context.Context, key string) (string, error)
}
