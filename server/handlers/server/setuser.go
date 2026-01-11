package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleSetUser(user *auth.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMISSIONS(commands.SETUSER, auth.ADMIN)
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.SETUSER, 2)
	}

	// username := args[0].Content
	// pass := args[1].Content

	// perms := args[2:]

	return resp.BYTE_OK, nil
}
