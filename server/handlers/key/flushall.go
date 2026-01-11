package key

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleFlushAll(user *auth.User, _ resp.ArgList) ([]byte, error) {
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.FLUSHALL, auth.WRITE)
	}

	internal.FlushAll()
	internal.FlushAllTTL()
	return resp.BYTE_OK, nil
}
