package key

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleFlushDB(user *authtypes.User, _ resp.ArgList) ([]byte, error) {
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.FLUSHALL, "WRITE")
	}

	user.DB.FlushAll()
	user.DB.FlushAllTTL()
	return resp.BYTE_OK, nil
}
