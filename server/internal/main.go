package internal

import (
	"sync"
)

var store = make(map[string]string)
var storeMu sync.RWMutex

func Set(key string, value string) {
	storeMu.Lock()
	defer storeMu.Unlock()

	store[key] = value
}

func Get(key string) string {
	storeMu.RLock()
	defer storeMu.RUnlock()

	return store[key]
}

func GetMany(keys []string) []string {
	out := make([]string, len(keys))
	storeMu.RLock()
	defer storeMu.RUnlock()

	for i, k := range keys {
		out[i] = store[k]
	}

	return out
}

func Del(key string) {
	storeMu.Lock()
	defer storeMu.Unlock()

	delete(store, key)
}

func FlushAll() {
	storeMu.Lock()
	defer storeMu.Unlock()

	store = make(map[string]string)
}
