package auth

import (
	"fmt"
	"mini-redis/server/auth/authtypes"
	"strings"

	"gopkg.in/yaml.v3"
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

func SetRules(username string, rules ...authtypes.Rule) ([]authtypes.User, error) {
	userFile, err := OpenACLFile(false)
	if err != nil {
		return nil, err
	}

	defer userFile.Close()

	decoder := yaml.NewDecoder(userFile)
	users := make([]authtypes.User, 0)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	found := false
	for _, u := range users {
		if u.Username == username {
			u.Rules.Add(rules)
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("user could not be found")
	}

	// Reopen the file in truncate mode for writing
	userFile, err = OpenACLFile(true)
	if err != nil {
		return nil, err
	}
	defer userFile.Close()

	encoder := yaml.NewEncoder(userFile)
	err = encoder.Encode(users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func RemoveRules(username string, rules ...authtypes.Rule) ([]authtypes.User, error) {
	userFile, err := OpenACLFile(false)
	if err != nil {
		return nil, err
	}

	defer userFile.Close()

	decoder := yaml.NewDecoder(userFile)
	users := make([]authtypes.User, 0)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	fmt.Println(users)

	found := false
	for _, u := range users {
		if u.Username == username {
			u.Rules.Subtract(rules)
			found = true
		}
	}

	fmt.Println(users)

	if !found {
		return nil, fmt.Errorf("user could not be found")
	}

	// Reopen the file in truncate mode for writing
	userFile, err = OpenACLFile(true)
	if err != nil {
		return nil, err
	}
	defer userFile.Close()

	encoder := yaml.NewEncoder(userFile)
	err = encoder.Encode(users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
