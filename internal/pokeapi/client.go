package pokeapi

import (
	"net/http"
	"time"

	"github.com/Tamelien/pokedex-cli/internal/pokecache"
)

const baseURLConst = "https://pokeapi.co/api/v2"

type Client struct {
	cache      *pokecache.Cache
	httpClient *http.Client
	baseURL    string
}

func NewClient(timeout, cacheInterval time.Duration) *Client {
	return &Client{
		cache:   pokecache.NewCache(cacheInterval),
		baseURL: baseURLConst,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}
