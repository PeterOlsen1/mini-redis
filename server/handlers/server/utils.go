package server

import (
	"log"
	"mini-redis/types/errors"
	"os"
	"path/filepath"
)

func getSaveFilePath(saveFile string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("LISTSAVES: failed to get home directory:", err)
		return "", errors.GENERAL
	}

	homeFolder := filepath.Join(homeDir, ".mini-redis")
	saveFilePath := filepath.Join(homeFolder, "backups", saveFile)
	return saveFilePath, nil
}

func getSaveFiles() ([]os.DirEntry, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("LISTSAVES: failed to get home directory:", err)
		return nil, errors.GENERAL
	}

	homeFolder := filepath.Join(homeDir, ".mini-redis")
	saveFolderPath := filepath.Join(homeFolder, "backups")
	err = os.MkdirAll(saveFolderPath, 0755)
	if err != nil {
		log.Println("LISTSAVES: failed to create .mini-redis/backups directory:", err)
		return nil, errors.GENERAL
	}

	files, err := os.ReadDir(saveFolderPath)
	if err != nil {
		log.Println("LISTSAVES: failed to list files in save folder:", err)
		return nil, errors.GENERAL
	}

	return files, nil
}

func getSaveFile(i int) (os.DirEntry, error) {
	files, err := getSaveFiles()
	if err != nil {
		return nil, err
	}

	if i < 0 {
		return files[0], nil
	}

	if i >= len(files) {
		return files[len(files)-1], nil
	}

	return files[i], nil
}
