package key

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
)

func HandleDel(args []resp.RESPItem) ([]byte, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("del requires 1 argument")
	}

	key := args[0].Content
	internal.Del(key)
	return resp.BYTE_OK, nil
}
