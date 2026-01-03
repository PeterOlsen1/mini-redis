package handlers

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/types"
	"strconv"
)

func handleSet(args []types.RESPItem) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("set requires 2 arguments")
	}

	key := args[0].Content
	value := args[1].Content

	// check if int or string
	num, err := strconv.Atoi(value)
	if err != nil {
		internal.Set(key, value, types.STRING)
	} else {
		internal.Set(key, num, types.INT)
	}

	return "OK", nil
}
