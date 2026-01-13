package server

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/internal"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"os"
	"path/filepath"
)

func HandleLoad(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMISSIONS(commands.LOAD, authtypes.ADMIN)
	}

	if len(args) < 1 {
		return nil, errors.ARG_COUNT(commands.LOAD, 1)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("failed to get home directory:", err)
		return nil, errors.GENERAL
	}

	homeFolder := filepath.Join(homeDir, ".mini-redis")
	loadFolderPath := filepath.Join(homeFolder, "backups")
	err = os.MkdirAll(loadFolderPath, 0755)
	if err != nil {
		fmt.Println("Failed to create .mini-redis/backups directory:", err)
		return nil, errors.GENERAL
	}

	loadFile, err := os.Open(filepath.Join(loadFolderPath, fmt.Sprintf("backup-%s.rdb", args.String(0))))
	if err != nil {
		fmt.Println("Failed to open backup file")
		return nil, errors.GENERAL
	}

	defer loadFile.Close()
	err = internal.Load(loadFile)
	if err != nil {
		return nil, errors.GENERAL
	}

	return resp.BYTE_OK, nil
}
