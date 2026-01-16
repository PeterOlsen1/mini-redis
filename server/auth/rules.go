package auth

import (
	"fmt"
	"mini-redis/server/auth/authtypes"
	"strings"
)

func ParseRules(rules ...string) authtypes.Ruleset {
	out := make([]authtypes.Rule, 0)
	for _, ruleString := range rules {
		// Ifs are mutually exclusive with admin last in the weird case that someone had a regex including "admin"
		if strings.Contains(ruleString, "read") {
			cut := strings.TrimSuffix(strings.TrimPrefix(ruleString, "read("), ")")
			readSplit := strings.SplitSeq(cut, " ")
			for readRule := range readSplit {
				mode := readRule[0]
				modeType := authtypes.ALLOW
				if mode == '-' {
					modeType = authtypes.DENY
				}

				rule := authtypes.Rule{
					Regex:     readRule[1:],
					Mode:      modeType,
					Operation: authtypes.READ,
				}
				out = append(out, rule)
			}
		} else if strings.Contains(ruleString, "write") {
			cut := strings.TrimSuffix(strings.TrimPrefix(ruleString, "write("), ")")
			writeSplit := strings.SplitSeq(cut, " ")
			for writeRule := range writeSplit {
				mode := writeRule[0]
				modeType := authtypes.ALLOW
				if mode == '-' {
					modeType = authtypes.DENY
				}

				rule := authtypes.Rule{
					Regex:     writeRule[1:],
					Mode:      modeType,
					Operation: authtypes.WRITE,
				}

				out = append(out, rule)
			}
		} else if strings.Contains(ruleString, "admin") {
			out = append(out, authtypes.ADMIN_RULE)
		}
	}

	return out
}

func AddRules(username string, rules authtypes.Ruleset) error {
	user, ok := GetUser(username)
	if !ok {
		return fmt.Errorf("user could not be found")
	}

	user.Rules.Add(rules)
	user.Perms = user.Rules.ExtractPerms()
	SetUser(user)

	return UpdateACLFile()
}

func RemoveRules(username string, rules authtypes.Ruleset) error {
	user, ok := GetUser(username)

	if !ok {
		return fmt.Errorf("user could not be found")
	}

	user.Rules.Subtract(rules)
	user.Perms = user.Rules.ExtractPerms()
	SetUser(user)

	return UpdateACLFile()
}
