package auth

import "strings"

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`

	// Perms is a bit array.
	// idx 0 = admin
	// idx 1 = read
	// idx 2 = write
	Perms int `yaml:"perms"`
}

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
		perms = append(perms, "admin")
	}

	if u.Read() {
		perms = append(perms, "read")
	}

	if u.Write() {
		perms = append(perms, "write")
	}

	if len(perms) == 0 {
		perms = append(perms, "none")
	}

	return strings.Join(perms, ", ")
}
