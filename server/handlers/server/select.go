package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func HandleSelect(user **authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.SELECT, 1)
	}

	dbNum, err := args.Int(0)
	if err != nil {
		return nil, errors.INVALID_ARG
	}

	db := internal.GetDB(dbNum)
	(*user).DB = db
	return resp.BYTE_OK, nil
}
