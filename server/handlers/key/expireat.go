package key

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"strconv"
)

func HandleExpireAt(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.EXPIREAT, 2)
	}

	key := args.String(0)
	if !user.CanWrite(key) {
		return nil, errors.PERMS_KEY(commands.EXPIREAT, authtypes.WRITE, key)
	}

	timeString := args[1].Content
	time, err := strconv.Atoi(timeString)
	if err != nil {
		return nil, fmt.Errorf("failed to convert time to integer")
	}

	return resp.BYTE_INT(internal.HandleExpireAt(key, time)), nil
}
