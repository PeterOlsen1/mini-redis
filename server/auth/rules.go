package auth

import (
	"fmt"
	"mini-redis/server/auth/authtypes"
	"strings"
)

func ParseRules(rules ...string) authtypes.Ruleset {
	out := make([]authtypes.Rule, len(rules))
	for i, ruleString := range rules {
		if strings.Contains(ruleString, "read") {
			cut := strings.TrimSuffix(strings.TrimPrefix(ruleString, "read("), ")")
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
			out[i] = rule
		}
		if strings.Contains(ruleString, "write") {
			cut := strings.TrimSuffix(strings.TrimPrefix(ruleString, "write("), ")")
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

			out[i] = rule
		}
	}

	return out
}

func AddRules(username string, rules ...authtypes.Rule) error {
	user, ok := GetUser(username)

	if !ok {
		return fmt.Errorf("user could not be found")
	}

	user.Rules.Add(rules)
	SetUser(user)

	return UpdateACLFile()
}

func RemoveRules(username string, rules ...authtypes.Rule) error {
	user, ok := GetUser(username)

	if !ok {
		return fmt.Errorf("user could not be found")
	}

	user.Rules.Subtract(rules)
	SetUser(user)

	return UpdateACLFile()
}
