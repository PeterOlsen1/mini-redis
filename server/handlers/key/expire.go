package key

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"strconv"
)

func HandleExpire(params []resp.RESPItem) ([]byte, error) {
	if len(params) < 2 {
		return nil, fmt.Errorf("expire requires 2 parameters")
	}

	key := params[0].Content
	ttl, err := strconv.Atoi(params[1].Content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TTL")
	}

	if internal.Get(key) == nil {
		return resp.BYTE_INT(0), nil
	}

	internal.SetTTL(key, ttl)
	return resp.BYTE_INT(1), nil
}
