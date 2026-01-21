package server

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"strings"
)

func HandleListSaves(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMISSIONS(commands.LISTSAVES, authtypes.ADMIN)
	}

	files, err := getSaveFiles()
	if err != nil {
		return nil, err
	}

	var out strings.Builder
	for i, f := range files {
		fmt.Fprintf(&out, "%d) %s\n", i, f.Name())
	}

	return resp.BYTE_STRING(out.String()), nil
}
