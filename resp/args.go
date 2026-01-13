package resp

import (
	"mini-redis/types/errors"
	"strconv"
	"strings"
)

type ArgList []RESPItem

func (l ArgList) Int(idx int) (int, error) {
	if len(l) <= idx {
		idx = len(l) - 1
	}

	item := l[idx]
	itemInt, err := strconv.Atoi(item.Content)
	if err != nil {
		return -1, errors.INVALID_ARG
	}

	return itemInt, nil
}

func (l ArgList) String(idx int) string {
	if len(l) <= idx {
		return l[len(l)-1].Content
	}

	if idx < 0 {
		return l[0].Content
	}

	item := l[idx]
	return item.Content
}

func (l ArgList) Slice(start int, end int) []string {
	if start < 0 {
		start = 0
	}

	if end > len(l) {
		end = len(l)
	}

	out := make([]string, end-start)
	j := 0
	for i := start; i < end; i++ {
		out[j] = l[i].Content
		j += 1
	}

	return out
}

func (l ArgList) Includes(substring string) bool {
	for _, item := range l {
		if strings.Contains(item.Content, substring) {
			return true
		}
	}

	return false
}

// Return the index of the first argument that contains the requested substring
func (l ArgList) SubstringIdx(substring string) int {
	for i, item := range l {
		if strings.Contains(item.Content, substring) {
			return i
		}
	}

	return -1
}
