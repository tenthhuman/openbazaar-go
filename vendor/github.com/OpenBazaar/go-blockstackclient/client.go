package blockstackclient

import (
	"encoding/json"
	"errors"
	"github.com/jbenet/go-multihash"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
	"time"
)

type BlockstackClient struct {
	resolverURL string
	cache       map[string]CachedGuid
	cacheLife   time.Duration
	sync.Mutex
}

type CachedGuid struct {
	guid   string
	exipry time.Time
}

func NewBlockStackClient(resolverURL string) *BlockstackClient {
	b := &BlockstackClient{
		resolverURL: resolverURL,
		cache:       make(map[string]CachedGuid),
		cacheLife:   time.Minute,
	}
	go b.gc()
	return b
}

func (b *BlockstackClient) Resolve(handle string) (guid string, err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			err = errors.New("Couldn't parse blockchainID json")
		}
	}()
	formatted := formatHandle(handle)
	if cached, ok := b.cache[formatted]; ok {
		return cached.guid, nil
	}
	resolver, err := url.Parse(b.resolverURL)
	if err != nil {
		return "", err
	}
	resolver.Path = path.Join(resolver.Path, "v2", "users", formatted)
	resp, err := http.Get(resolver.String())
	if err != nil {
		return "", errors.New("Error querying resolver")
	}
	decoder := json.NewDecoder(resp.Body)
	var data map[string]interface{}
	err = decoder.Decode(&data)
	if err != nil {
		return "", err
	}
	obj := data[formatted].(map[string]interface{})
	profile := obj["profile"].(map[string]interface{})
	account := profile["account"].([]interface{})
	for _, a := range account {
		acc := a.(map[string]interface{})
		service := strings.ToLower(acc["service"].(string))
		identifier := acc["identifier"].(string)
		if service == "openbazaar" {
			mh, err := multihash.FromB58String(identifier)
			if err != nil {
				return "", err
			}
			_, err = multihash.Decode(mh)
			if err != nil {
				return "", err
			}
			b.Lock()
			b.cache[formatted] = CachedGuid{identifier, time.Now().Add(b.cacheLife)}
			b.Unlock()
			return identifier, nil
		}
	}
	return "", errors.New("Handle does not exist")
}

func (b *BlockstackClient) SetCacheLifetime(duration time.Duration) {
	b.cacheLife = duration
}

func (b *BlockstackClient) gc() {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			b.deleteExpiredCache()
		}
	}
}

func (b *BlockstackClient) deleteExpiredCache() {
	b.Lock()
	defer b.Unlock()
	for k, v := range b.cache {
		if v.exipry.Before(time.Now()) {
			delete(b.cache, k)
		}
	}
}

func formatHandle(handle string) string {
	if handle[0:1] == "@" {
		return strings.ToLower(handle[1:])
	}
	return strings.ToLower(handle)
}
