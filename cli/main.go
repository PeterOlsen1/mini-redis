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
		os.Exit(1)
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
		if handleLineIn(c, input) != nil {
			break
		}
	}
}

func handleLineIn(c *client.RedisClient, input string) error {
	tokens := strings.Split(input, " ")

	if len(input) >= 2 && input[0] == '!' {
		if input[1] != '!' {
			execNumString := strings.TrimPrefix(input, "!")
			execNum, err := strconv.Atoi(execNumString)
			if err != nil || execNum < 0 {
				fmt.Printf("History command must be a positive integer")
				return nil
			}

			execHistory(c, execNum+1)
		} else {
			execHistory(c, 1)
		}

		return nil
	}
	if tokens[0] == "history" {
		if len(tokens) == 1 {
			showHistory(10)
			return nil
		}

		historyLen, err := strconv.Atoi(tokens[1])
		if err != nil || historyLen < 0 {
			fmt.Println("History length must be a positive integer")
			return nil
		}

		showHistory(historyLen)
		return nil
	}

	if input == "exit" {
		fmt.Println("Exiting...")
		return fmt.Errorf("")
	}

	resp, err := handleInput(c, tokens)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

	return nil
}

func handleInput(c *client.RedisClient, tokens []string) (string, error) {
	if len(tokens) < 1 {
		return "", fmt.Errorf("too few input tokens")
	}

	if commands.ParseCommand(tokens[0]) == commands.NONE {
		return "", fmt.Errorf("Invalid command.")
	}

	cmd := commands.ParseCommand(tokens[0])
	req := client.InitRequest(cmd)
	for i := 1; i < len(tokens); i++ {
		req.AddParam(tokens[i])
	}

	// need extra buffer space for info command
	if cmd == commands.INFO {
		req.SetBufSize(4096)
	}

	return c.SendAndReceive(req)
}
