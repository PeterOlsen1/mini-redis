package key

import (
	"mini-redis/resp"
	"mini-redis/server/internal"
)

func HandleFlushAll(_ []resp.RESPItem) (string, error) {
	internal.FlushAll()
	internal.FlushAllTTL()
	return "OK", nil
}
