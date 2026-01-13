package server

import (
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
	"strings"
)

func HandleSetUser(user *authtypes.User, args resp.ArgList) ([]byte, error) {
	if !user.Admin() {
		return nil, errors.PERMISSIONS(commands.SETUSER, authtypes.ADMIN)
	}

	if len(args) < 2 {
		return nil, errors.ARG_COUNT(commands.SETUSER, 2)
	}

	username := args[0].Content
	pass := args[1].Content

	rules := make([]authtypes.Rule, 0)
	perms := 0
	if args.Includes("admin") {
		perms |= authtypes.ADMIN
	}
	if idx := args.SubstringIdx("read"); idx != -1 {
		if strings.Contains(args.String(idx), "(") {
			cut := strings.TrimSuffix(strings.TrimPrefix(args.String(idx), "read("), ")")
			mode := cut[0]
			modeType := authtypes.ALLOW
			if mode == '-' {
				modeType = authtypes.DENY
			}

			rule := authtypes.Rule{
				Regex:     cut[1:],
				Mode:      modeType,
				Operation: authtypes.READ,
			}
			rules = append(rules, rule)
		}
		perms |= authtypes.READ
	}
	if idx := args.SubstringIdx("write"); idx != -1 {
		if strings.Contains(args.String(idx), "(") {
			cut := strings.TrimSuffix(strings.TrimPrefix(args.String(idx), "write("), ")")
			mode := cut[0]
			modeType := authtypes.ALLOW
			if mode == '-' {
				modeType = authtypes.DENY
			}

			rule := authtypes.Rule{
				Regex:     cut[1:],
				Mode:      modeType,
				Operation: authtypes.WRITE,
			}
			rules = append(rules, rule)
		}
		perms |= authtypes.WRITE
	}

	err := auth.AddACLUser(username, pass, perms, rules)
	if err != nil {
		return nil, err
	}
	return resp.BYTE_OK, nil
}
