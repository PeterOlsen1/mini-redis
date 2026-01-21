package key

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"strconv"
)

func HandleExpire(user *authtypes.User, params resp.ArgList) ([]byte, error) {
	if len(params) < 2 {
		return nil, errors.ARG_COUNT(commands.EXPIRE, 2)
	}

	key := params[0].Content
	if !user.CanWrite(key) {
		return nil, errors.PERMS_KEY(commands.EXPIRE, "WRITE", key)
	}

	ttl, err := strconv.Atoi(params[1].Content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TTL")
	}

	if user.DB.Get(key) == nil {
		return resp.BYTE_INT(0), nil
	}

	user.DB.SetTTL(key, ttl)
	return resp.BYTE_INT(1), nil
}
