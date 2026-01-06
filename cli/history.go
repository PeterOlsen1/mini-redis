package main

import (
	"bufio"
	"fmt"
	"mini-redis/client"
	"os"
	"path/filepath"
)

var historyFile *os.File
var numLines = 0

func countHistoryLines() int {
	historyFile.Seek(0, 0)
	scanner := bufio.NewScanner(historyFile)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func openHistoryFile() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}

	homeFolder := filepath.Join(homeDir, ".mini-redis")
	historyFilePath := filepath.Join(homeFolder, "history")

	err = os.MkdirAll(homeFolder, 0755)
	if err != nil {
		fmt.Println("Failed to create .mini-redis directory:", err)
		os.Exit(1)
	}

	historyFile, err = os.OpenFile(historyFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Failed to open or create history file:", err)
		os.Exit(1)
	}

	numLines = countHistoryLines()
}

func writeHistory(line string) {
	if historyFile == nil {
		fmt.Printf("failed to write history")
		os.Exit(1)
	}

	_, err := historyFile.WriteString(line + "\n")
	if err != nil {
		fmt.Printf("failed to write history")
		os.Exit(1)
	}

	numLines += 1
}

func showHistory(n int) {
	if historyFile == nil || n < 0 {
		fmt.Printf("Failed to show history")
		return
	}

	n = min(n, numLines)

	historyFile.Seek(0, 0)
	scanner := bufio.NewScanner(historyFile)

	lines := make([]string, 0, n)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) > n {
			lines = lines[1:]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading history file:", err)
		return
	}

	for i, line := range lines {
		fmt.Printf("%d %s\n", len(lines)-(i+1), line)
	}
}

func execHistory(c *client.RedisClient, n int) {
	historyFile.Seek(0, 0)
	scanner := bufio.NewScanner(historyFile)

	n = min(n, numLines)
	line := ""
	for range numLines - n {
		scanner.Scan()
		line = scanner.Text()
	}

	// skip infinite recursion case
	if line[0] == '!' {
		fmt.Println("Cannot repeat history command")
		return
	}

	fmt.Println(line)
	handleLineIn(c, line)
}
