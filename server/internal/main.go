package internal

import (
	"fmt"
	"mini-redis/types"
	"mini-redis/types/errors"
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

func Set(key string, value any) {
	storeMu.Lock()
	defer storeMu.Unlock()

	store[key] = newItem(value, types.STRING)
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
		store[key] = newItem("1", types.STRING)
		return "1", true
	}

	if store[key].Type != types.STRING {
		return "0", false
	}

	oldVal, err := strconv.Atoi(store[key].Item.(string))
	if err != nil {
		return "0", false
	}

	newVal := strconv.Itoa(oldVal + 1)
	store[key].Item = newVal
	return newVal, true
}

func Decr(key string) (string, bool) {
	storeMu.Lock()
	defer storeMu.Unlock()

	if store[key] == nil {
		store[key] = newItem("-1", types.STRING)
		return "-1", true
	}

	if store[key].Type != types.STRING {
		return "0", false
	}

	oldVal, err := strconv.Atoi(store[key].Item.(string))
	if err != nil {
		return "0", false
	}

	newVal := strconv.Itoa(oldVal - 1)
	store[key].Item = newVal
	return newVal, true
}

func LPush(key string, values []string) int {
	storeMu.Lock()
	defer storeMu.Unlock()

	if store[key] == nil {
		new := newItem(values, types.ARRAY)
		store[key] = new

		return len(values)
	}

	items, ok := store[key].Item.([]string)
	if !ok {
		return -1
	}

	// append LEFT
	items = append(values, items...)
	return len(items)
}

func RPush(key string, values []string) int {
	storeMu.Lock()
	defer storeMu.Unlock()

	if store[key] == nil {
		new := newItem(values, types.ARRAY)
		store[key] = new

		return len(values)
	}

	items, ok := store[key].Item.([]string)
	if !ok {
		return -1
	}

	// append RIGHT
	items = append(items, values...)
	return len(items)
}

func LPop(key string, num int) ([]string, error) {
	storeMu.Lock()
	defer storeMu.Unlock()

	val := store[key]
	if val == nil {
		return nil, nil
	}
	if val.Type != types.ARRAY {
		return nil, fmt.Errorf(errors.WRONGTYPE)
	}

	arr := val.Array()
	if num <= 0 {
		return []string{}, nil
	}
	if num >= len(arr) {
		delete(store, key)
		return arr, nil
	}

	ret := arr[:num]
	val.Item = arr[num:]

	return ret, nil
}

func RPop(key string, num int) ([]string, error) {
	storeMu.Lock()
	defer storeMu.Unlock()

	val := store[key]
	if val == nil {
		return nil, nil
	}
	if val.Type != types.ARRAY {
		return nil, fmt.Errorf(errors.WRONGTYPE)
	}

	arr := val.Array()
	if num <= 0 {
		return []string{}, nil
	}
	if num >= len(arr) {
		delete(store, key)
		return arr, nil
	}

	ret := arr[num:]
	val.Item = arr[:num]

	return ret, nil
}
