package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleAddRule(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMS_GENERAL(commands.ADDRULE)
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.ADDRULE, 2)
	}

	dstUser := args.String(0)
	rules := auth.ParseRules(args.Slice(1, 100000)...)

	err := auth.AddRules(dstUser, rules)
	if err != nil {
		return nil, errors.GENERAL
	}

	// async ACL file update
	go auth.UpdateACLFile()

	return resp.BYTE_OK, nil
}
