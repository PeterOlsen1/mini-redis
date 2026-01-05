package key

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
)

func HandleTTL(params []resp.RESPItem) ([]byte, error) {
	if len(params) < 1 {
		return nil, fmt.Errorf("TTL requires 1 parameter")
	}

	// return -2 on non-existent key
	key := params[0].Content
	if internal.Get(key) == nil {
		return resp.BYTE_INT(-2), nil
	}

	// no associated TTL
	ttl := internal.GetTTL(key)
	if ttl == -1 {
		return resp.BYTE_INT(-1), nil
	}

	// ttl expired
	if ttl == -2 {
		internal.Del(key)
		internal.DelTTL(key)
		return resp.BYTE_INT(-2), nil
	}

	return resp.BYTE_INT(ttl), nil
}
