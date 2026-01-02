package handlers

import (
	"mini-redis/server/internal"
	"mini-redis/types"
)

func handleFlushAll(_ []types.RESPItem) (string, error) {
	internal.FlushAll()
	return "OK", nil
}
