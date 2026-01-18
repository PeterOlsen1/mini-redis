package key

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleTTL(user *authtypes.User, params resp.ArgList) ([]byte, error) {
	if len(params) < 1 {
		return nil, errors.ARG_COUNT(commands.TTL, 1)
	}

	// return -2 on non-existent key
	key := params.String(0)
	if !user.CanRead(key) {
		return nil, errors.PERMS_KEY(commands.TTL, authtypes.READ, key)
	}

	if internal.Get(key) == nil {
		return resp.BYTE_INT(-2), nil
	}

	// no associated TTL
	ttl := internal.GetTTL(key)
	if ttl == -1 {
		return resp.BYTE_INT(-1), nil
	}

	// ttl expired
	if ttl == -2 {
		internal.Del(key)
		internal.DelTTL(key)
		return resp.BYTE_INT(-2), nil
	}

	return resp.BYTE_INT(ttl), nil
}
