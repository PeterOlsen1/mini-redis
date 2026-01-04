package key

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"strconv"
)

func HandleExists(args []resp.RESPItem) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("get requires 1 argument")
	}

	stringArgs := make([]string, len(args))
	for i, a := range args {
		stringArgs[i] = a.Content
	}
	results := internal.GetMany(stringArgs)

	count := 0
	for _, r := range results {
		if r != "" {
			count += 1
		}
	}
	return strconv.Itoa(count), nil
}
