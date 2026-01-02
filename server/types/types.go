package types

type RESPItem struct {
	Len     int
	Content string
	Command Command
}

type Command int

const (
	NONE Command = iota
	PING
	ECHO
	SET
	GET
	DEL
	EXISTS
	EXPIRE
	TTL
	INCR
	LPUSH
	RPUSH
	LPOP
	RPOP
	INFO
)
