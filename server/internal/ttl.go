package internal

import (
	"sync"
	"time"
)

// Lazily check TTL values instead of using a gorotine for overhead
// When

var expiration = make(map[string]int64)
var ttlMu sync.RWMutex

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
