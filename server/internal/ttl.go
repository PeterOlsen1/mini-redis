package internal

import (
	"time"
)

// func StartTTLScan(ctx context.Context) {
// 	if cfg.Server.TTLCheck <= 0 {
// 		return
// 	}

// 	ticker := time.NewTicker(time.Duration(cfg.Server.TTLCheck) * time.Millisecond)

// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case <-ticker.C:
// 				CheckTTLs()
// 			}
// 		}
// 	}()
// }

func (db *Database) CheckTTLs() {
	db.ttlMu.Lock()
	defer db.ttlMu.Unlock()

	for k := range db.ttlStore {
		if db.ttlStore[k]-time.Now().UnixMilli() <= 0 {
			delete(db.ttlStore, k)

			db.mu.Lock()
			delete(db.store, k)
			db.mu.Unlock()
			// info.Expire()
		}
	}
}

func (db *Database) SetTTL(key string, seconds int) {
	db.ttlMu.Lock()
	defer db.ttlMu.Unlock()

	exp := time.Now().UnixMilli() + int64(seconds*1000)
	db.ttlStore[key] = exp
}

// Returns -1 on no TTL and -2 on expired
func (db *Database) GetTTL(key string) int {
	db.ttlMu.RLock()
	exp, ok := db.ttlStore[key]
	db.ttlMu.RUnlock()
	if !ok {
		return -1
	}

	ttl := int(exp-time.Now().UnixMilli()) / 1000
	if ttl <= 0 {
		db.DelTTL(key)
		return -2
	}
	return ttl
}

func (db *Database) DelTTL(key string) {
	db.ttlMu.Lock()
	defer db.ttlMu.Unlock()

	delete(db.ttlStore, key)
}

func (db *Database) FlushAllTTL() {
	db.ttlMu.Lock()
	defer db.ttlMu.Unlock()

	db.ttlStore = make(map[string]int64)
}

func (db *Database) HandleExpireAt(key string, secs int) int {
	curSecs := int(time.Now().UnixMilli() / 1000)

	db.mu.RLock()
	_, ok := db.store[key]
	db.mu.RUnlock()
	if !ok {
		return 0
	}

	if secs < curSecs {
		db.mu.Lock()
		delete(db.store, key)
		db.mu.Unlock()

		db.ttlMu.Lock()
		delete(db.ttlStore, key)
		db.ttlMu.Unlock()
		// info.Expire()
		return 0
	}

	db.ttlMu.Lock()
	defer db.ttlMu.Unlock()

	db.ttlStore[key] = int64(secs)
	return 1
}

func (db *Database) HandleExpireTime(key string) int64 {
	db.mu.RLock()
	_, ok := db.store[key]
	if !ok {
		db.mu.RUnlock() // FORGOT UNLOCK HERE AND DEADLOCKED!!!!!
		return -2
	}
	db.mu.RUnlock()

	db.ttlMu.RLock()
	defer db.ttlMu.RUnlock()

	time, ok := db.ttlStore[key]
	if !ok {
		return -1
	}

	return time
}
