package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/cfg"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleAuth(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if user.Username != "" && user.Perms == 0 {
		return nil, errors.ALREADY_AUTH
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.AUTH, 2)
	}

	username := args[0].Content
	pass := args[1].Content

	for _, u := range cfg.Server.Users {
		if u.Username == username && u.Password == pass {
			user.Username = u.Username
			user.Perms = u.Perms
			return resp.BYTE_OK, nil
		}
	}

	newUser, err := auth.CheckACLUser(username, pass)
	if err != nil {
		return nil, errors.COULD_NOT_AUTHENTICATE
	}

	if newUser != nil {
		user = newUser
		return resp.BYTE_OK, nil
	}

	return nil, errors.COULD_NOT_AUTHENTICATE
}
