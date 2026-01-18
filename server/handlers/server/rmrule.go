package server

import (
	"log"
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleRMRule(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMS_GENERAL(commands.RMRULE)
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.RMRULE, 2)
	}

	dstUser := args.String(0)
	rules := auth.ParseRules(args.Slice(1, 100000)...)

	err := auth.RemoveRules(dstUser, rules)
	if err != nil {
		log.Println(err)
		return nil, errors.GENERAL
	}

	// update ACL file in goroutine
	go auth.UpdateACLFile()

	return resp.BYTE_OK, nil
}
