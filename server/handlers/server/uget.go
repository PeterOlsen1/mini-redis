package server

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/cfg"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"strings"
)

func HandleUGet(user *auth.User, args resp.ArgList) ([]byte, error) {
	if len(args) == 0 && !user.Admin() {
		return nil, errors.PERMISSIONS(commands.UGET, auth.ADMIN)
	}

	// no args means list all users
	if len(args) == 0 && user.Admin() {
		var out strings.Builder
		out.WriteString("Defined users:\n")
		for i, u := range cfg.Server.Users {
			fmt.Fprintf(&out, "%d) %s: %s\n", i+1, u.Username, u.PermString())
		}

		out.WriteString("\nACL users:\n")
		for i, u := range cfg.Server.LoadedUsers {
			fmt.Fprintf(&out, "%d) %s: %s\n", i+1, u.Username, u.PermString())
		}

		return resp.BYTE_STRING(out.String()), nil
	}

	if user.Username != args.String(0) && !user.Admin() {
		return nil, errors.PERMS_GENERAL(commands.UGET)
	}

	requested := args.String(0)
	for _, u := range cfg.Server.LoadedUsers {
		if u.Username == requested {
			return resp.BYTE_STRING(fmt.Sprintf("user: %s\npermissions: %s\n", args.String(0), u.PermString())), nil
		}
	}

	return nil, errors.GENERAL
}
