package resp

import (
	"fmt"
	"strconv"
)

type ArgList []RESPItem

func (l ArgList) Int(idx int) (int, error) {
	if len(l) <= idx {
		return -1, fmt.Errorf("arg index is out of range")
	}

	item := l[idx]
	itemInt, err := strconv.Atoi(item.Content)
	if err != nil {
		return -1, err
	}

	return itemInt, nil
}

func (l ArgList) String(idx int) (string, error) {
	if len(l) <= idx {
		return "", fmt.Errorf("arg index is out of range")
	}

	item := l[idx]
	return item.Content, nil
}
