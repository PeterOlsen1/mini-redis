package key

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleKeys(user auth.User, _ []resp.RESPItem) ([]byte, error) {
	if !user.Read() {
		return nil, errors.PERMISSIONS(commands.KEYS, auth.READ)
	}

	serialized, err := resp.Serialize(internal.Keys(), resp.ARRAY)
	if err != nil {
		return nil, err
	}
	return serialized, nil
}
