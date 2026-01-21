package internal

import (
	"maps"
	"mini-redis/types/errors"
	"strconv"
	"sync"
)

type Database struct {
	store map[string]*StoreItem
	mu    sync.RWMutex

	ttlStore map[string]int64
	ttlMu    sync.RWMutex

	Number int
}

func NewDb(number int) *Database {
	return &Database{
		store:  make(map[string]*StoreItem),
		Number: number,
	}
}

func (db *Database) Size() int {
	db.mu.RLock()
	defer db.mu.RUnlock()

	return len(db.store)
}

func (db *Database) Set(key string, value any) {
	// info.SetOp()
	db.mu.Lock()
	db.store[key] = newItem(value, STRING)
	db.mu.Unlock()

	db.DelTTL(key)
}

func (db *Database) Get(key string) *StoreItem {
	ttl := db.GetTTL(key)
	// info.GetOp()

	db.mu.RLock()
	defer db.mu.RUnlock()

	if ttl == -2 {
		// delete if TTL is expired
		// info.Expire()
		delete(db.store, key)
		return nil
	}

	return db.store[key]
}

func (db *Database) GetMany(keys []string) []any {
	out := make([]any, len(keys))
	// info.GetOp()
	db.mu.RLock()
	defer db.mu.RUnlock()

	for i, key := range keys {
		ttl := db.GetTTL(key) // TTL expired
		if ttl == -2 {
			// unsafe delete, don't call methods becuase of locking
			delete(db.store, key)
			out[i] = ""
		} else {
			out[i] = db.store[key].Item
		}
	}

	return out
}

// Not necessary to delete TTL, as if a new key with the same name is added,
// it is deleted there
func (db *Database) Del(key string) {
	// info.Delete()
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.store, key)
}

func (db *Database) FlushAll() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.store = make(map[string]*StoreItem)
}

func (db *Database) Incr(key string) (string, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.store[key] == nil {
		db.store[key] = newItem("1", STRING)
		return "1", true
	}

	if db.store[key].Type != STRING {
		return "0", false
	}

	oldVal, err := strconv.Atoi(db.store[key].Item.(string))
	if err != nil {
		return "0", false
	}

	newVal := strconv.Itoa(oldVal + 1)
	db.store[key].Item = newVal
	return newVal, true
}

func (db *Database) Decr(key string) (string, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.store[key] == nil {
		db.store[key] = newItem("-1", STRING)
		return "-1", true
	}

	if db.store[key].Type != STRING {
		return "0", false
	}

	oldVal, err := strconv.Atoi(db.store[key].Item.(string))
	if err != nil {
		return "0", false
	}

	newVal := strconv.Itoa(oldVal - 1)
	db.store[key].Item = newVal
	return newVal, true
}

func (db *Database) LPush(key string, values []string) int {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.store[key] == nil {
		new := newItem(values, ARRAY)
		db.store[key] = new

		return len(values)
	}

	items, ok := db.store[key].Item.([]string)
	if !ok {
		return -1
	}

	// append LEFT
	items = append(values, items...)
	db.store[key].Item = items

	return len(items)
}

func (db *Database) RPush(key string, values []string) int {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.store[key] == nil {
		new := newItem(values, ARRAY)
		db.store[key] = new

		return len(values)
	}

	items, ok := db.store[key].Item.([]string)
	if !ok {
		return -1
	}

	// append RIGHT
	items = append(items, values...)
	db.store[key].Item = items

	return len(items)
}

func (db *Database) LPop(key string, num int) ([]string, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	val := db.store[key]
	if val == nil {
		return nil, nil
	}
	if val.Type != ARRAY {
		// return nil, errors.WRONGTYPE
	}

	arr, err := val.Array()
	if err != nil {
		return nil, err
	}

	if num <= 0 {
		return []string{}, nil
	}
	if num >= len(arr) {
		delete(db.store, key)
		return arr, nil
	}

	ret := arr[:num]
	val.Item = arr[num:]

	return ret, nil
}

func (db *Database) RPop(key string, num int) ([]string, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	val := db.store[key]
	if val == nil {
		return nil, nil
	}
	if val.Type != ARRAY {
		return nil, errors.WRONGTYPE
	}

	if num <= 0 {
		return []string{}, nil
	}

	arr, err := val.Array()
	if err != nil {
		return nil, err
	}

	if num >= len(arr) {
		delete(db.store, key)
		return arr, nil
	}

	ret := arr[len(arr)-num:]
	val.Item = arr[:len(arr)-num]

	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return ret, nil
}

func (db *Database) LRange(key string, start int, end int) ([]string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	empty := []string{}
	val := db.store[key]
	if val == nil {
		return empty, nil
	}

	if val.Type != ARRAY {
		// return nil, errors.WRONGTYPE
	}

	arr, ok := val.Item.([]string)
	if !ok {
		// return nil, errors.WRONGTYPE
	}

	if end < start {
		return empty, nil
	}

	if start > len(arr) {
		return empty, nil
	}

	if start < 0 {
		start = 0
	}

	if end > len(arr) {
		end = len(arr) - 1
	}

	if start == end {
		return []string{arr[start]}, nil
	}

	return arr[start:][:end+1], nil
}

func (db *Database) LGet(key string) ([]string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	val := db.store[key]
	if val == nil {
		return nil, nil
	}

	if val.Type != ARRAY {
		// return nil, errors.WRONGTYPE
	}

	arr, ok := val.Item.([]string)
	if !ok {
		// return nil, errors.WRONGTYPE
	}

	return arr, nil
}

func (db *Database) Keys() []string {
	db.mu.RLock()
	keys := maps.Keys(db.store)
	db.mu.RUnlock()

	out := make([]string, 0)
	for k := range keys {
		out = append(out, k)
	}

	return out
}
