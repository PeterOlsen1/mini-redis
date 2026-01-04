package list

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/internal"
	"strconv"
	"strings"
)

func HandleRPop(args []resp.RESPItem) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("LPOP requires 1 arguments")
	}

	key := args[0].Content
	if len(args) == 1 {
		res, err := internal.RPop(key, 1)
		if err != nil {
			return "", err
		}

		if len(res) == 0 {
			return "", fmt.Errorf("cannot pop empty list")
		}

		return res[0], nil
	}

	num, err := strconv.Atoi(args[1].Content)
	if err != nil {
		return "", fmt.Errorf("invalid pop count")
	}
	res, err := internal.RPop(key, num)
	if err != nil {
		return "", err
	}

	// make this more complex later
	return strings.Join(res, ","), nil
}
