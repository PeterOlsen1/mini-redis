package main

import (
	"bufio"
	"fmt"
	"mini-redis/client"
	"mini-redis/types/commands"
	"os"
	"strconv"
	"strings"
)

func main() {
	c, err := client.NewClient(nil)
	if err != nil {
		fmt.Println("failed to establish redis connection, exiting...")
		os.Exit(-1)
	}

	openHistoryFile()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Try again")
			continue
		}

		input = strings.TrimSpace(input)
		writeHistory(input)
		tokens := strings.Split(input, " ")

		if tokens[0] == "history" {
			if len(tokens) == 1 {
				showHistory(10)
				continue
			}

			historyLen, err := strconv.Atoi(tokens[1])
			if err != nil || historyLen < 0 {
				fmt.Println("History length must be a positive integer")
				continue
			}

			showHistory(historyLen)
			continue
		}

		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		resp, err := handleInput(c, tokens)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	}
}

func handleInput(c *client.RedisClient, tokens []string) (string, error) {
	fmt.Println("Input:", tokens)

	if len(tokens) < 1 {
		return "", fmt.Errorf("too few input tokens")
	}

	if commands.ParseCommand(tokens[0]) == commands.NONE {
		return "", fmt.Errorf("Invalid command.")
	}

	req := client.InitRequest(tokens[0])
	for i := 1; i < len(tokens); i++ {
		req.AddParam(tokens[i])
	}

	return c.SendAndReceive(req)
}
