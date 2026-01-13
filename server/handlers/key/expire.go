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

func HandleExpire(user *authtypes.User, params resp.ArgList) ([]byte, error) {
	if !user.Write() {
		return nil, errors.PERMISSIONS(commands.EXPIRE, authtypes.WRITE)
	}

	if len(params) < 2 {
		return nil, errors.ARG_COUNT(commands.EXPIRE, 2)
	}

	key := params[0].Content
	ttl, err := strconv.Atoi(params[1].Content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TTL")
	}

	if internal.Get(key) == nil {
		return resp.BYTE_INT(0), nil
	}

	internal.SetTTL(key, ttl)
	return resp.BYTE_INT(1), nil
}
