package key

import (
	"mini-redis/resp"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleExpireTime(args []resp.RESPItem) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.EXPIRETIME, 1)
	}

	key := args[0].Content
	ret := internal.HandleExpireTime(key)

	return resp.BYTE_INT(int(ret)), nil
}
