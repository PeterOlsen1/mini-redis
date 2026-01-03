package list

import (
	"fmt"
	"mini-redis/server/internal"
	"mini-redis/types"
	"strconv"
	"strings"
)

func HandleLPop(args []types.RESPItem) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("LPOP requires 1 arguments")
	}

	key := args[0].Content
	if len(args) == 1 {
		res, err := internal.LPop(key, 1)

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
	res, err := internal.LPop(key, num)
	if err != nil {
		return "", err
	}

	// make this more complex later
	return strings.Join(res, ","), nil
}
