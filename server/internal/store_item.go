package internal

import (
	"mini-redis/types/errors"
)

type StoreItem struct {
	Type StoreType
	Item any
}

func (i *StoreItem) Array() ([]string, error) {
	arr, ok := i.Item.([]string)
	if !ok {
		return nil, errors.WRONGTYPE
	}
	return arr, nil
}

func (i *StoreItem) String() string {
	return i.Item.(string)
}

type StoreType int

const (
	STRING StoreType = iota
	ARRAY
)

func newItem(value any, storeType StoreType) *StoreItem {
	return &StoreItem{
		Item: value,
		Type: storeType,
	}
}
