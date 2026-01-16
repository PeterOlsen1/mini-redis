package authtypes

import (
	"regexp"
	"strings"
)

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`

	// Perms is a bit array.
	// idx 0 = admin
	// idx 1 = read
	// idx 2 = write
	Perms int `yaml:"perms"`

	Rules Ruleset `yaml:"rules"`
}

type UserPermission int

const ADMIN = 0b1
const READ = 0b10
const WRITE = 0b100

var authRequired = false

func SetAuthRequired(new bool) {
	authRequired = new
}

func (u User) Admin() bool {
	if !authRequired {
		return true
	}

	return u.Perms&ADMIN != 0
}

func (u User) Read() bool {
	if !authRequired {
		return true
	}

	return u.Perms&READ != 0
}

func (u User) Write() bool {
	if !authRequired {
		return true
	}

	return u.Perms&WRITE != 0
}

func (u User) PermString() string {
	perms := make([]string, 0)
	if u.Admin() {
		perms = append(perms, "ADMIN")
	}

	if u.Read() {
		perms = append(perms, "READ")
	}

	if u.Write() {
		perms = append(perms, "WRITE")
	}

	if len(perms) == 0 {
		perms = append(perms, "NONE")
	}

	return strings.Join(perms, ", ")
}

func (u User) CanRead(key string) bool {
	if u.Admin() {
		return true
	}

	// general read permission
	// check out negatives
	if u.Read() {
		for _, rule := range u.Rules.Negatives().Read() {
			matched, err := regexp.Match(rule.Regex, []byte(key))
			if err == nil && matched == true {
				return false
			}
		}

		return true
	}

	for _, rule := range u.Rules {
		if rule.Operation != READ {
			continue
		}

		matched, err := regexp.Match(rule.Regex, []byte(key))
		if err != nil || matched == false {
			continue
		}

		return true
	}

	return false
}

func (u User) CanWrite(key string) bool {
	if u.Admin() {
		return true
	}

	// general write permission
	if u.Write() {
		for _, rule := range u.Rules.Negatives().Write() {
			matched, err := regexp.Match(rule.Regex, []byte(key))
			if err == nil && matched == true {
				return false
			}
		}

		return true
	}

	for _, rule := range u.Rules {
		if rule.Operation != WRITE {
			continue
		}

		matched, err := regexp.Match(rule.Regex, []byte(key))
		if err != nil || matched == false {
			continue
		}

		return true
	}

	return false
}
