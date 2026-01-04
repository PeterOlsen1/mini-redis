package key

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"strconv"
)

func HandleExpire(params []resp.RESPItem) (string, error) {
	if len(params) < 2 {
		return "", fmt.Errorf("expire requires 2 parameters")
	}

	key := params[0].Content
	ttl, err := strconv.Atoi(params[1].Content)
	if err != nil {
		return "", fmt.Errorf("failed to parse TTL")
	}

	if internal.Get(key) == nil {
		return "0", nil
	}

	internal.SetTTL(key, ttl)
	return "1", nil
}
