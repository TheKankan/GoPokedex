package pokeapi

import (
	"time"

	"github.com/TheKankan/GoPokedex/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func NewCache(cacheInterval time.Duration) pokecache.Cache {
	return *pokecache.NewCache(cacheInterval)
}
