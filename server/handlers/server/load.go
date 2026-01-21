package server

import (
	"log"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"os"
)

func HandleLoad(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMISSIONS(commands.LOAD, authtypes.ADMIN)
	}

	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.LOAD, 1)
	}

	fileIdx, err := args.Int(0)
	if err != nil {
		return nil, err
	}

	fileName, err := getSaveFile(fileIdx)
	if err != nil {
		return nil, errors.GENERAL
	}

	loadFilePath, err := getSaveFilePath(fileName.Name())
	if err != nil {
		return nil, errors.GENERAL
	}

	loadFile, err := os.Open(loadFilePath)
	if err != nil {
		log.Println("Failed to open backup file")
		return nil, errors.GENERAL
	}

	defer loadFile.Close()
	err = internal.Load(loadFile)
	if err != nil {
		return nil, errors.GENERAL
	}

	return resp.BYTE_OK, nil
}
