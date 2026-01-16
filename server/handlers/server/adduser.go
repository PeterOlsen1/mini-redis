package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleAddUser(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMISSIONS(commands.ADDUSER, authtypes.ADMIN)
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.ADDUSER, 2)
	}

	username := args[0].Content
	pass := args[1].Content

	perms := 0
	if args.Includes("admin") {
		perms |= authtypes.ADMIN
	}

	err := auth.AddACLUser(&user, username, pass)
	if err != nil {
		return nil, err
	}

	ruleSlice := args.Slice(2, 10000)
	user.Rules = auth.ParseRules(ruleSlice...)
	user.Perms |= user.Rules.ExtractPerms()
	return resp.BYTE_OK, nil
}
