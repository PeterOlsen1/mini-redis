package key

import (
	"mini-redis/resp"
	"mini-redis/server/internal"
)

func HandleKeys(_ []resp.RESPItem) ([]byte, error) {
	serialized, err := resp.Serialize(internal.Keys(), resp.ARRAY)
	if err != nil {
		return nil, err
	}
	return serialized, nil
}
