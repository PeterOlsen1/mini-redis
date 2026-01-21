package internal

import (
	"encoding/gob"
	"os"
	"sync"
)

type StoreItem struct {
	Type StoreType
	Item any
}

func (i *StoreItem) Array() []string {
	return i.Item.([]string)
}
func (i *StoreItem) String() string {
	return i.Item.(string)
}

type StoreType int

const (
	STRING StoreType = iota
	ARRAY
)

var store = make(map[int]*Database)
var storeMu sync.RWMutex

func InitStore(numDBs int) {
	for i := range numDBs {
		store[i] = NewDb()
	}
}

func GetStoreSize() int {
	storeMu.RLock()
	defer storeMu.RUnlock()

	return len(store)
}

func newItem(value any, storeType StoreType) *StoreItem {
	return &StoreItem{
		Item: value,
		Type: storeType,
	}
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
