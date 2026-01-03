package types

type RESPItem struct {
	Len     int
	Content string
}

type StoreItem struct {
	Type StoreType
	Item any
}

type StoreType int

const (
	INT StoreType = iota
	STRING
	ARRAY
)
