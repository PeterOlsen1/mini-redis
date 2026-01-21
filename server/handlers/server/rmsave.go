package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"os"
)

func HandleRMSave(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMISSIONS(commands.RMSAVE, "ADMIN")
	}

	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.RMSAVE, 1)
	}

	fileIdx, err := args.Int(0)
	if err != nil {
		return nil, errors.INVALID_ARG
	}

	file, err := getSaveFile(fileIdx)
	if err != nil {
		return nil, errors.GENERAL
	}

	filePath, err := getSaveFilePath(file.Name())
	if err != nil {
		return nil, errors.GENERAL
	}

	if err = os.Remove(filePath); err != nil {
		return nil, errors.GENERAL
	}

	return resp.BYTE_OK, nil
}
