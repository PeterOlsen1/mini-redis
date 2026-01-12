package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/cfg"
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

	username := args[0].Content
	pass := args[1].Content

	admin := args.Includes("admin")
	read := args.Includes("read")
	write := args.Includes("write")

	perms := 0
	if admin {
		perms |= auth.ADMIN
	}
	if read {
		perms |= auth.READ
	}
	if write {
		perms |= auth.WRITE
	}

	users, err := auth.AddACLUser(username, pass, perms)
	if err != nil {
		return nil, err
	}

	// update loaded user list. Can't be done in AddACLUser bc circular import
	cfg.Server.LoadedUsers = users
	return resp.BYTE_OK, nil
}
