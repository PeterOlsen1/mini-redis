package server

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/cfg"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"strings"
)

func HandleUGet(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if len(args) == 0 && !user.Admin() {
		return nil, errors.PERMISSIONS(commands.UGET, "ADMIN")
	}

	// no args means list all users
	if len(args) == 0 && user.Admin() {
		var out strings.Builder
		out.WriteString("Config defined users:\n")
		for i, u := range cfg.Server.Users {
			fmt.Fprintf(&out, "%d) %s: %s\n", i, u.Username, u.PermString())
		}

		out.WriteString("\nACL users:\n")
		for i, u := range auth.GetAllUsers() {
			fmt.Fprintf(&out, "%d) %s: %s\n", i, u.Username, u.PermString())

			if len(u.Rules) > 0 {
				fmt.Fprintf(&out, "   User rules:\n")

				for _, rule := range u.Rules {
					operationString := "READ"
					if rule.Operation == authtypes.WRITE {
						operationString = "WRITE"
					} else if rule.Operation == authtypes.ADMIN {
						continue
					}

					if rule.Mode {
						fmt.Fprintf(&out, "   + on %s: %s\n", operationString, rule.Regex)
					} else {
						fmt.Fprintf(&out, "   - on %s: %s\n", operationString, rule.Regex)
					}
				}
			}
		}

		return resp.BYTE_STRING(out.String()), nil
	}

	if user.Username != args.String(0) && !user.Admin() {
		return nil, errors.PERMS_GENERAL(commands.UGET)
	}

	requested := args.String(0)
	for _, u := range auth.GetAllUsers() {
		if u.Username == requested {
			var out strings.Builder
			fmt.Fprintf(&out, "%s: %s\n", u.Username, u.PermString())

			if len(u.Rules) > 0 {
				fmt.Fprintf(&out, "   User rules:\n")

				for _, rule := range u.Rules {
					operationString := "READ"
					if rule.Operation == authtypes.WRITE {
						operationString = "WRITE"
					}

					if rule.Mode {
						fmt.Fprintf(&out, "   + on %s: %s\n", operationString, rule.Regex)
					} else {
						fmt.Fprintf(&out, "   - on %s: %s\n", operationString, rule.Regex)
					}
				}
			}

			return resp.BYTE_STRING(out.String()), nil
		}
	}

	return nil, errors.GENERAL
}
