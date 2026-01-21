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
	"time"
)

func HandleSave(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMISSIONS(commands.SAVE, "ADMIN")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("failed to get home directory:", err)
		return nil, errors.GENERAL
	}

	homeFolder := filepath.Join(homeDir, ".mini-redis")
	saveFolderPath := filepath.Join(homeFolder, "backups")
	err = os.MkdirAll(saveFolderPath, 0755)
	if err != nil {
		fmt.Println("Failed to create .mini-redis/backups directory:", err)
		return nil, errors.GENERAL
	}

	today := time.Now().Format(time.RFC3339)
	saveFile, err := os.Create(filepath.Join(saveFolderPath, fmt.Sprintf("backup-%s.rdb", today)))
	if err != nil {
		fmt.Println("Failed to create backup file")
		return nil, errors.GENERAL
	}

	defer saveFile.Close()
	err = internal.Save(saveFile)
	if err != nil {
		return nil, errors.GENERAL
	}

	return resp.BYTE_OK, nil
}
