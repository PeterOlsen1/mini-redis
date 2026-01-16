package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/cfg"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleAuth(userPtr **authtypes.User, args resp.ArgList) ([]byte, error) {
	user := *userPtr
	if user.Username != "" && user.Perms == 0 {
		return nil, errors.ALREADY_AUTH
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.AUTH, 2)
	}

	username := args.String(0)
	pass := args.String(1)

	for _, u := range cfg.Server.Users {
		if u.Username == username && u.Password == pass {
			user.Username = u.Username
			user.Perms = u.Perms
			return resp.BYTE_OK, nil
		}
	}

	err := auth.CheckACLUser(userPtr, username, pass)
	if err != nil {
		return nil, errors.COULD_NOT_AUTHENTICATE
	}

	return resp.BYTE_OK, nil
}
