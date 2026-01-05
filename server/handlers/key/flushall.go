package key

import (
	"mini-redis/resp"
	"mini-redis/server/internal"
)

func HandleFlushAll(_ []resp.RESPItem) ([]byte, error) {
	internal.FlushAll()
	internal.FlushAllTTL()
	return resp.OK, nil
}
