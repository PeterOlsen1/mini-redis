package key

import (
	"mini-redis/server/internal"
	"mini-redis/types"
)

func HandleFlushAll(_ []types.RESPItem) (string, error) {
	internal.FlushAll()
	internal.FlushAllTTL()
	return "OK", nil
}
