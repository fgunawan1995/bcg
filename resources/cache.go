package resources

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

const (
	cacheExpiration = 24 * time.Hour
	cacheCleanup    = 5 * time.Minute
)

//InitCache create new cache
func InitCache() *cache.Cache {
	return cache.New(cacheExpiration, cacheCleanup)
}
