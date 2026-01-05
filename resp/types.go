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

var respTypeStrings = [...]string{
	"STRING",
	"ERROR",
	"ARRAY",
	"BULK_STRING",
	"NULL",
}

func (t RespType) String() string {
	return respTypeStrings[t]
}
