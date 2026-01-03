package internal

import (
	"mini-redis/types"
	"strconv"
	"sync"
)

var store = make(map[string]*types.StoreItem)
var storeMu sync.RWMutex

func newItem(value any, storeType types.StoreType) *types.StoreItem {
	return &types.StoreItem{
		Item: value,
		Type: storeType,
	}
}

func Set(key string, value any, storeType types.StoreType) {
	storeMu.Lock()
	defer storeMu.Unlock()

	store[key] = newItem(value, storeType)
	DelTTL(key)
}

func Get(key string) *types.StoreItem {
	ttl := GetTTL(key)
	if ttl == -2 {
		// key has TTL and it has expired
		// deletion of the TTL entry is handled by GetTTL
		Del(key)
		return nil
	}

	storeMu.RLock()
	defer storeMu.RUnlock()

	return store[key]
}

func GetMany(keys []string) []any {
	out := make([]any, len(keys))
	storeMu.RLock()
	defer storeMu.RUnlock()

	for i, key := range keys {
		ttl := GetTTL(key) // TTL expired
		if ttl == -2 {
			// unsafe delete, don't call methods becuase of locking
			delete(store, key)
			out[i] = ""
		} else {
			out[i] = store[key].Item
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

	store = make(map[string]*types.StoreItem)
}

func Incr(key string) (string, bool) {
	storeMu.Lock()
	defer storeMu.Unlock()

	if store[key] == nil {
		store[key] = newItem(1, types.INT)
		return "1", true
	}

	if store[key].Type != types.INT {
		return "0", false
	}

	oldVal := store[key].Item.(int)
	store[key].Item = oldVal + 1
	return strconv.Itoa(oldVal + 1), true
}

func Decr(key string) (string, bool) {
	storeMu.Lock()
	defer storeMu.Unlock()

	if store[key] == nil {
		store[key] = newItem(-1, types.INT)
		return "-1", true
	}

	if store[key].Type != types.INT {
		return "0", false
	}

	oldVal := store[key].Item.(int)
	store[key].Item = oldVal - 1
	return strconv.Itoa(oldVal - 1), true
}
