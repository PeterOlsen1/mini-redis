package key

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleExists(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Read() {
		return nil, errors.PERMISSIONS(commands.EXISTS, authtypes.READ)
	}

	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.EXISTS, 1)
	}

	stringArgs := make([]string, len(args))
	for i, a := range args {
		stringArgs[i] = a.Content
	}
	results := internal.GetMany(stringArgs)

	count := 0
	for _, r := range results {
		if r != "" {
			count += 1
		}
	}
	return resp.BYTE_INT(count), nil
}
