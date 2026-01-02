package internal

import (
	"sync"
)

var store = make(map[string]string)
var mu sync.RWMutex

func Set(key string, value string) {
	mu.Lock()
	defer mu.Unlock()

	store[key] = value
}

func Get(key string) string {
	mu.RLock()
	defer mu.RUnlock()

	return store[key]
}

func Del(key string) {
	mu.Lock()
	defer mu.Unlock()

	delete(store, key)
}

func FlushAll() {
	mu.Lock()
	defer mu.Unlock()

	store = make(map[string]string)
}
