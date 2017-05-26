package util

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

var c = cache.New(4*time.Hour, 8*time.Hour)

func GetFromCache(key string) (interface{}, bool) {
	return c.Get(key)
}

func PutInCache(key string, data interface{}) {
	c.SetDefault(key, data)
}
