package main

import (
	"sync"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var (
	cache *gocache.Cache
	once  sync.Once
)

func Init() {
	if cache == nil {
		once.Do(func() {
			cache = gocache.New(5*time.Minute, 10*time.Minute)
		})
	}
}

func GetCacheClient() *gocache.Cache {
	if cache == nil {
		Init()
	}
	return cache
}
