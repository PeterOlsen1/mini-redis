package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/cfg"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"strings"
)

func HandleSetUser(user *auth.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMISSIONS(commands.SETUSER, auth.ADMIN)
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.SETUSER, 2)
	}

	username := args[0].Content
	pass := args[1].Content

	rules := make([]auth.Rule, 0)
	perms := 0
	if args.Includes("admin") {
		perms |= auth.ADMIN
	}
	if idx := args.SubstringIdx("read"); idx != -1 {
		if strings.Contains(args.String(idx), "(") {
			cut := strings.TrimSuffix(strings.TrimPrefix(args.String(idx), "("), ")")
			mode := cut[0]
			modeType := auth.ALLOW
			if mode == '-' {
				modeType = auth.DENY
			}

			rule := auth.Rule{
				Regex:     cut[1:],
				Mode:      modeType,
				Operation: auth.READ,
			}
			rules = append(rules, rule)
		}
		perms |= auth.READ
	}
	if idx := args.SubstringIdx("write"); idx != -1 {
		if strings.Contains(args.String(idx), "(") {
			cut := strings.TrimSuffix(strings.TrimPrefix(args.String(idx), "("), ")")
			mode := cut[0]
			modeType := auth.ALLOW
			if mode == '-' {
				modeType = auth.DENY
			}

			rule := auth.Rule{
				Regex:     cut[1:],
				Mode:      modeType,
				Operation: auth.WRITE,
			}
			rules = append(rules, rule)
		}
		perms |= auth.WRITE
	}

	users, err := auth.AddACLUser(username, pass, perms, rules)
	if err != nil {
		return nil, err
	}

	// update loaded user list. Can't be done in AddACLUser bc circular import
	cfg.Server.LoadedUsers = users
	return resp.BYTE_OK, nil
}
