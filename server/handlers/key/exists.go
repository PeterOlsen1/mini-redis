package key

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleExists(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.EXISTS, 1)
	}

	stringArgs := make([]string, len(args))
	for i, a := range args {
		if !user.CanRead(a.Content) {
			return nil, errors.PERMS_KEY(commands.EXISTS, "READ", a.Content)
		}

		stringArgs[i] = a.Content
	}
	results := user.DB.GetMany(stringArgs)

	count := 0
	for _, r := range results {
		if r != "" {
			count += 1
		}
	}
	return resp.BYTE_INT(count), nil
}
