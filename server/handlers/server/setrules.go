package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/cfg"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleSetRules(user *auth.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMS_GENERAL(commands.SETRULE)
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.SETRULE, 2)
	}

	dstUser := args.String(0)
	rules := auth.ParseRules(args.Slice(1, 100000)...)

	newUsers, err := auth.SetRules(dstUser, rules...)
	if err != nil {
		return nil, errors.GENERAL
	}
	cfg.Server.LoadedUsers = newUsers

	return resp.BYTE_OK, nil
}
