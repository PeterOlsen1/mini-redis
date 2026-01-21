package internal

import (
	"encoding/gob"
	"os"
	"sync"
)

var store = make(map[int]*Database)
var storeMu sync.RWMutex

func InitStore(numDBs int) {
	for i := range numDBs {
		store[i] = NewDb(i)
	}
}

func GetDB(dbNum int) *Database {
	storeMu.RLock()
	defer storeMu.RUnlock()

	return store[dbNum]
}

func GetStoreSize() int {
	storeMu.RLock()
	defer storeMu.RUnlock()

	return len(store)
}

func Save(file *os.File) error {
	storeMu.RLock()
	defer storeMu.RUnlock()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(store)
}

func Load(file *os.File) error {
	storeMu.Lock()
	defer storeMu.Unlock()

	store = make(map[int]*Database)
	decoder := gob.NewDecoder(file)
	return decoder.Decode(&store)
}
