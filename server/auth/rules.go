package auth

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Rule struct {
	// The regex matching which keys the user can access
	Regex string `yaml:"regex"`

	// Either ALLOW or DENY
	Mode bool `yaml:"mode"`

	// The operation this rule is set on. Either READ or WRITE, not admin
	Operation UserPermission `yaml:"operation"`
}

const ALLOW = true
const DENY = false

type Ruleset []Rule

func (rset Ruleset) Contains(rule Rule) bool {
	for _, r := range rset {
		if r.Regex == rule.Regex && r.Mode == rule.Mode && r.Operation == rule.Operation {
			return true
		}
	}

	return false
}

func (rset Ruleset) Remove(rule Rule) {
	for i, r := range rset {
		if r.Regex == rule.Regex && r.Mode == rule.Mode && r.Operation == rule.Operation {
			rset = append(rset[:i], rset[i+1:]...)
			return
		}
	}
}

// Combines two rule sets. No identical items will be included
func (rset Ruleset) Add(other Ruleset) {
	for _, rule := range other {
		if rset.Contains(rule) {
			continue
		}

		rset = append(rset, rule)
	}
}

func (rset Ruleset) Subtract(other Ruleset) {
	for _, rule := range other {
		rset.Remove(rule)
	}
}

func ParseRules(rules ...string) Ruleset {
	out := make([]Rule, len(rules))
	for i, ruleString := range rules {
		if strings.Contains(ruleString, "read") {
			cut := strings.TrimSuffix(strings.TrimPrefix(ruleString, "read("), ")")
			mode := cut[0]
			modeType := ALLOW
			if mode == '-' {
				modeType = DENY
			}

			rule := Rule{
				Regex:     cut[1:],
				Mode:      modeType,
				Operation: READ,
			}
			out[i] = rule
		}
		if strings.Contains(ruleString, "write") {
			cut := strings.TrimSuffix(strings.TrimPrefix(ruleString, "write("), ")")
			mode := cut[0]
			modeType := ALLOW
			if mode == '-' {
				modeType = DENY
			}

			rule := Rule{
				Regex:     cut[1:],
				Mode:      modeType,
				Operation: WRITE,
			}

			out[i] = rule
		}
	}

	return out
}

func SetRules(username string, rules ...Rule) ([]User, error) {
	userFile, err := OpenACLFile(false)
	if err != nil {
		return nil, err
	}

	defer userFile.Close()

	decoder := yaml.NewDecoder(userFile)
	users := make([]User, 0)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	found := false
	for _, u := range users {
		if u.Username == username {
			u.Rules = rules
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

func RemoveRules(username string, rules ...Rule) ([]User, error) {
	userFile, err := OpenACLFile(false)
	if err != nil {
		return nil, err
	}

	defer userFile.Close()

	decoder := yaml.NewDecoder(userFile)
	users := make([]User, 0)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	found := false
	for _, u := range users {
		if u.Username == username {
			u.Rules.Subtract(rules)
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
