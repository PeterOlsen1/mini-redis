package key

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"strconv"
)

func HandleTTL(params []resp.RESPItem) (string, error) {
	if len(params) < 1 {
		return "", fmt.Errorf("TTL requires 1 parameter")
	}

	// return -2 on non-existent key
	key := params[0].Content
	if internal.Get(key) == nil {
		return "-2", nil
	}

	// no associated TTL
	ttl := internal.GetTTL(key)
	if ttl == -1 {
		return "-1", nil
	}

	// ttl expired
	if ttl == -2 {
		internal.Del(key)
		internal.DelTTL(key)
		return "-2", nil
	}

	return strconv.Itoa(ttl), nil
}
