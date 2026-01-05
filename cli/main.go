package main

import (
	"bufio"
	"fmt"
	"mini-redis/client"
	"mini-redis/types/commands"
	"os"
	"strings"
)

func main() {
	c, err := client.NewClient(nil)
	if err != nil {
		fmt.Println("failed to establish redis connection, exiting...")
		os.Exit(-1)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Try again")
			continue
		}

		input = strings.TrimSpace(input)
		tokens := strings.Split(input, " ")
		fmt.Println("Input:", tokens)

		resp, err := handleInput(c, tokens)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}

		if input == "exit" {
			fmt.Println("Exiting...")
			break
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
