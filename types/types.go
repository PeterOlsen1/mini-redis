package types

import (
	"net"
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

type Connection struct {
	Conn net.Conn
	Auth bool
}
