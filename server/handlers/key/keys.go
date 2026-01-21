package key

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleKeys(user *authtypes.User, _ resp.ArgList) ([]byte, error) {
	if !user.Read() {
		return nil, errors.PERMISSIONS(commands.KEYS, "READ")
	}

	serialized, err := resp.Serialize(user.DB.Keys(), resp.ARRAY)
	if err != nil {
		return nil, err
	}
	return serialized, nil
}
