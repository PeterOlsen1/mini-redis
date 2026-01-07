package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var curFile *os.File

func StartLogger() error {
	ticker := time.NewTicker(time.Millisecond * 500)

	_, err := makeNewFile()
	if err != nil {
		return err
	}

	go func() {
		for range ticker.C {
			upgrade, err := readFileSize()
			if err != nil {
				continue
			}

			if upgrade {
				curFile.Close()
				makeNewFile()
			}
		}
	}()

	return nil
}

// returns current log file size
func readFileSize() (bool, error) {
	info, err := curFile.Stat()
	if err != nil {
		return false, err
	}

	size := info.Size()
	if size > 1<<20 {
		return true, nil
	}

	return false, nil
}

func makeNewFile() (*os.File, error) {
	homeDir, _ := os.UserHomeDir()
	logDir := filepath.Join(homeDir, ".mini-redis", "logs")
	os.MkdirAll(logDir, os.ModePerm)

	time := time.Now().Format(time.RFC3339)
	outFile, err := os.Open(filepath.Join(logDir, fmt.Sprintf("mini-redis-%s.log", time)))
	if err != nil {
		return nil, err
	}

	curFile = outFile
	return outFile, nil
}
