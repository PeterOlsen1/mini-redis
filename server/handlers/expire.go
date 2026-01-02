package handlers

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/types"
	"strconv"
)

func handleExpire(params []types.RESPItem) (string, error) {
	if len(params) < 2 {
		return "", fmt.Errorf("expire requires 2 parameters")
	}

	key := params[0].Content
	ttl, err := strconv.Atoi(params[1].Content)
	if err != nil {
		return "", fmt.Errorf("failed to parse TTL")
	}

	if internal.Get(key) == "" {
		return "0", nil
	}

	internal.SetTTL(key, ttl)
	return "1", nil
}
