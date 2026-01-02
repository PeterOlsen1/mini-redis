package internal

import (
	"strconv"
	"sync"
)

var store = make(map[string]string)
var storeMu sync.RWMutex

func Set(key string, value string) {
	storeMu.Lock()
	defer storeMu.Unlock()

	store[key] = value
	DelTTL(key)
}

func Get(key string) string {
	ttl := GetTTL(key)
	if ttl == -2 {
		// key has TTL and it has expired
		// deletion of the TTL entry is handled by GetTTL
		Del(key)
		return ""
	}

	storeMu.RLock()
	defer storeMu.RUnlock()

	return store[key]
}

func GetMany(keys []string) []string {
	out := make([]string, len(keys))
	storeMu.RLock()
	defer storeMu.RUnlock()

	for i, key := range keys {
		ttl := GetTTL(key) // TTL expired
		if ttl == -2 {
			// unsafe delete, don't call methods becuase of locking
			delete(store, key)
			out[i] = ""
		} else {
			out[i] = store[key]
		}
	}

	return out
}

// Not necessary to delete TTL, as if a new key with the same name is added,
// it is deleted there
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

func Incr(key string) (string, bool) {
	storeMu.Lock()
	defer storeMu.Unlock()

	if store[key] == "" {
		store[key] = "1"
		return "1", true
	}

	val, err := strconv.Atoi(store[key])
	if err != nil {
		return "0", false
	}

	newStr := strconv.Itoa(val + 1)
	store[key] = newStr
	return newStr, true
}

func Decr(key string) (string, bool) {
	storeMu.Lock()
	defer storeMu.Unlock()

	if store[key] == "" {
		store[key] = "-1"
		return "-1", true
	}

	val, err := strconv.Atoi(store[key])
	if err != nil {
		return "0", false
	}

	newStr := strconv.Itoa(val - 1)
	store[key] = newStr
	return newStr, true
}
