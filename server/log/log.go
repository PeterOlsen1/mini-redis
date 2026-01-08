package log

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var curFile *os.File

func StartLogger(ctx context.Context) error {
	ticker := time.NewTicker(time.Second * 3) // 3 seconds, slow for testing

	_, err := makeNewFile()
	if err != nil {
		fmt.Printf("error creating log file: %e", err)
		return err
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				upgrade, err := readFileSize()
				if err != nil {
					continue
				}

				if upgrade {
					curFile.Close()
					makeNewFile()
				}
			case <-ctx.Done():
				curFile.Close()
				return
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
	outFile, err := os.Create(filepath.Join(logDir, fmt.Sprintf("%s.log", time)))
	if err != nil {
		return nil, err
	}

	log.SetOutput(outFile)
	curFile = outFile
	return outFile, nil
}
