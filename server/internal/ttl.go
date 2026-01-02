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

func GetTTL(key string) int {
	ttlMu.RLock()
	defer ttlMu.RUnlock()

	if expiration[key] < time.Now().UnixMilli() {
		return -1
	}

	return int(expiration[key]-time.Now().UnixMilli()) / 1000
}
