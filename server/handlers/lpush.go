package handlers

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/types"
	"strconv"
)

func handleLPush(args []types.RESPItem) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("LPUSH requires 2 arguments")
	}

	key := args[0].Content
	vals := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		vals[i] = arg.Content
	}

	ret := internal.LPush(key, vals)
	if ret == -1 {
		return "", fmt.Errorf("Operation against a key holding the wrong kind of value")
	}

	return strconv.Itoa(ret), nil
}
