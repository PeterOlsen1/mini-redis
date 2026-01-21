package authtypes

import (
	"mini-redis/server/internal"
	"regexp"
	"strings"
	"sync"
)

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`

	// Perms is a bit array.
	// idx 0 = admin
	// idx 1 = read
	// idx 2 = write
	Perms int `yaml:"perms"`

	// Set of rules for the user's operations
	Rules Ruleset `yaml:"rules"`

	// The current database that the user is connected to
	DB *internal.Database

	// Mutex for locking
	mu sync.Mutex
}

type UserPermission int

const ADMIN = 0b1
const READ = 0b10
const WRITE = 0b100

func (p UserPermission) String() string {
	switch p {
	case ADMIN:
		return "ADMIN"
	case READ:
		return "READ"
	case WRITE:
		return "WRITE"
	}

	return ""
}

var authRequired = true

func NewUser(username string, password string) User {
	return User{
		Username: username,
		Password: password,
		Perms:    0,
		Rules:    make(Ruleset, 0),
	}
}

func SetAuthRequired(new bool) {
	authRequired = new
}

func (u *User) SetPerms(new int) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.Perms = new
}

// rule functions

func (u *User) SubtractRules(rules Ruleset) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.Rules.Subtract(rules)
	u.Perms = u.Rules.ExtractPerms()
}

func (u *User) AddRules(rules Ruleset) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.Rules.Add(rules)
	u.Perms = u.Rules.ExtractPerms()
}

// permission getters

func (u *User) NoAuth() bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	return u.Password == "" && u.Username == ""
}

func (u *User) Admin() bool {
	if !authRequired {
		return true
	}
	u.mu.Lock()
	defer u.mu.Unlock()

	return u.Perms&ADMIN != 0
}

func (u *User) Read() bool {
	if !authRequired {
		return true
	}
	u.mu.Lock()
	defer u.mu.Unlock()

	return u.Perms&READ != 0
}

func (u *User) Write() bool {
	if !authRequired {
		return true
	}
	u.mu.Lock()
	defer u.mu.Unlock()

	return u.Perms&WRITE != 0
}

func (u *User) PermString() string {
	u.mu.Lock()
	defer u.mu.Unlock()

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

func (u *User) CanRead(key string) bool {
	if u.Admin() {
		return true
	}
	u.mu.Lock()
	defer u.mu.Unlock()

	// Check negatives
	for _, rule := range u.Rules.Iter().Negatives().Read() {
		matched, err := regexp.Match(rule.Regex, []byte(key))
		if err == nil && matched {
			return false
		}
	}

	// Check general read bit
	if u.Perms&READ != 0 {
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

func (u *User) CanWrite(key string) bool {
	if u.Admin() {
		return true
	}
	u.mu.Lock()
	defer u.mu.Unlock()

	// Check negatives
	for _, rule := range u.Rules.Iter().Negatives().Write() {
		matched, err := regexp.Match(rule.Regex, []byte(key))
		if err == nil && matched {
			return false
		}
	}

	// Check general write bit
	if u.Perms&WRITE != 0 {
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
