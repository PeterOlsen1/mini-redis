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
		return nil, errors.PERMISSIONS(commands.ADDUSER, "ADMIN")
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.ADDUSER, 2)
	}

	username := args.String(0)
	pass := args.String(1)

	if testUsr, exists := auth.GetUser(username); exists == true || testUsr != nil {
		return nil, errors.USER_EXISTS
	}

	newUser, err := auth.AddACLUser(username, pass)
	if err != nil {
		return nil, err
	}

	ruleSlice := args.Slice(2, 10000)
	newUser.AddRules(auth.ParseRules(ruleSlice...))

	// update the ACL file
	go auth.UpdateACLFile()

	return resp.BYTE_OK, nil
}
