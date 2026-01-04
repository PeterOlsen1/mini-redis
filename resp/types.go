package resp

type RESPItem struct {
	Len     int
	Content string
}

type RespType int

const (
	STRING RespType = iota
	ERR
	ARRAY
	BULK_STRING
	NULL
)
