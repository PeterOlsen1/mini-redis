package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/cfg"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleRMUser(user *auth.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.RMUSER, 1)
	}

	delUser := args.String(0)

	if user == nil || (user.Username != delUser && !user.Admin()) {
		return nil, errors.PERMS_GENERAL(commands.RMUSER)
	}

	newUsers, err := auth.RemoveACLUser(delUser)
	if err != nil {
		return nil, err
	}

	cfg.Server.LoadedUsers = newUsers

	if user.Username == delUser {
		user.Username = ""
		user.Perms = 0
	}

	return resp.BYTE_OK, nil
}
