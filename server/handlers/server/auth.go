package server

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/cfg"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleAuth(user *auth.User, args resp.ArgList) ([]byte, error) {
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
		return nil, fmt.Errorf("could not authenticate")
	}

	if newUser != nil {
		user.Username = newUser.Username
		user.Perms = newUser.Perms
		// don't set the password bc we don't need that
		return resp.BYTE_OK, nil
	}

	return nil, fmt.Errorf("could not authenticate")
}
