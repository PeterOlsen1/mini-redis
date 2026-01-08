package internal

import (
	"context"
	"mini-redis/server/cfg"
	"mini-redis/server/info"
	"sync"
	"time"
)

// Lazily check TTL values instead of using a gorotine for overhead
var expiration = make(map[string]int64)
var ttlMu sync.RWMutex

func StartTTLScan(ctx context.Context) {
	if cfg.Server.TTLCheck <= 0 {
		return
	}

	ticker := time.NewTicker(time.Duration(cfg.Server.TTLCheck) * time.Millisecond)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				CheckTTLs()
			}
		}
	}()
}

func CheckTTLs() {
	ttlMu.Lock()
	defer ttlMu.Unlock()

	for k := range expiration {
		if expiration[k]-time.Now().UnixMilli() <= 0 {
			delete(expiration, k)

			storeMu.Lock()
			delete(store, k)
			storeMu.Unlock()
			info.Expire()
		}
	}
}

func SetTTL(key string, seconds int) {
	ttlMu.Lock()
	defer ttlMu.Unlock()

	exp := time.Now().UnixMilli() + int64(seconds*1000)
	expiration[key] = exp
}

// Returns -1 on no TTL and -2 on expired
func GetTTL(key string) int {
	ttlMu.RLock()
	exp, ok := expiration[key]
	ttlMu.RUnlock()
	if !ok {
		return -1
	}

	ttl := int(exp-time.Now().UnixMilli()) / 1000
	if ttl <= 0 {
		DelTTL(key)
		return -2
	}
	return ttl
}

func DelTTL(key string) {
	ttlMu.Lock()
	defer ttlMu.Unlock()

	delete(expiration, key)
}

func FlushAllTTL() {
	ttlMu.Lock()
	defer ttlMu.Unlock()

	expiration = make(map[string]int64)
}

func HandleExpireAt(key string, secs int) int {
	curSecs := int(time.Now().UnixMilli() / 1000)

	storeMu.RLock()
	_, ok := store[key]
	storeMu.RUnlock()
	if !ok {
		return 0
	}

	if secs < curSecs {
		storeMu.Lock()
		delete(store, key)
		storeMu.Unlock()

		ttlMu.Lock()
		delete(expiration, key)
		ttlMu.Unlock()
		info.Expire()
		return 0
	}

	ttlMu.Lock()
	defer ttlMu.Unlock()

	expiration[key] = int64(secs)
	return 1
}

func HandleExpireTime(key string) int64 {
	storeMu.RLock()
	_, ok := store[key]
	if !ok {
		return -2
	}
	storeMu.RUnlock()

	ttlMu.RLock()
	defer ttlMu.RUnlock()

	time, ok := expiration[key]
	if !ok {
		return -1
	}

	return time
}
