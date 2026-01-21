package server

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"strings"
)

func HandleCommand(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) == 0 {
		var out strings.Builder
		for i := range commands.Len() {
			if i == 0 {
				continue
			}

			cmd := commands.Command(i)
			fmt.Fprintf(&out, "%d) %s - %s\n", i, cmd.String(), cmd.Info())
		}

		fmt.Fprintf(&out, "%d commands total.\n", commands.Len()-1)
		return resp.BYTE_STRING(out.String()), nil
	}

	cmd := commands.ParseCommand(args.String(0))
	if cmd == commands.NONE {
		return nil, errors.INVALID_ARG
	}

	return resp.BYTE_STRING(fmt.Sprintf("%s - %s\n", cmd.String(), cmd.Info())), nil
}
