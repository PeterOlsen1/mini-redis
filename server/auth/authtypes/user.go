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

var authRequired = true

func SetAuthRequired(new bool) {
	authRequired = new
}

func (u User) NoAuth() bool {
	return u.Password == "" && u.Username == ""
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
	if u.Perms&ADMIN != 0 {
		perms = append(perms, "ADMIN")
	}

	if u.Perms&READ != 0 {
		perms = append(perms, "READ")
	}

	if u.Perms&WRITE != 0 {
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

	// Check negatives
	for _, rule := range u.Rules.Iter().Negatives().Read() {
		matched, err := regexp.Match(rule.Regex, []byte(key))
		if err == nil && matched {
			return false
		}
	}

	// Check general read bit
	if u.Read() {
		return true
	}

	// Check positives
	for _, rule := range u.Rules.Iter().Positives().Read() {
		matched, err := regexp.Match(rule.Regex, []byte(key))
		if err == nil && matched {
			return true
		}
	}

	// Default to false
	return false
}

func (u User) CanWrite(key string) bool {
	if u.Admin() {
		return true
	}

	// Check negatives
	for _, rule := range u.Rules.Iter().Negatives().Write() {
		matched, err := regexp.Match(rule.Regex, []byte(key))
		if err == nil && matched {
			return false
		}
	}

	// Check general write bit
	if u.Write() {
		return true
	}

	// Check positives
	for _, rule := range u.Rules.Iter().Positives().Write() {
		matched, err := regexp.Match(rule.Regex, []byte(key))
		if err == nil && matched {
			return true
		}
	}

	// Default to false
	return false
}
