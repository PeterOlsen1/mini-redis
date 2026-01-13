package errors

import (
	"fmt"
	"mini-redis/server/auth"
	"mini-redis/types/commands"
)

var WRONGTYPE = fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
var NOT_INTEGER = fmt.Errorf("value is not an integer or out of range")
var INVALID_ARG = fmt.Errorf("INVALID_ARG failed to parse integer argument") //make into a function to return invaid argument?
func ARG_COUNT(cmd commands.Command, count int) error {
	return fmt.Errorf("%s requires %d arguments", cmd, count)
}

func PERMS_GENERAL(cmd commands.Command) error {
	return fmt.Errorf("you do not have permissions to run %s", cmd.String())
}

func PERMISSIONS(cmd commands.Command, perm int) error {
	switch perm {
	case auth.ADMIN:
		return fmt.Errorf("%s requires admin privileges", cmd.String())
	case auth.READ:
		return fmt.Errorf("%s requires read privileges", cmd.String())
	case auth.WRITE:
		return fmt.Errorf("%s requires write privileges", cmd.String())
	}

	return nil
}

var ALREADY_AUTH = fmt.Errorf("user is already authenticated. use logout to destroy session")
var COULD_NOT_AUTHENTICATE = fmt.Errorf("could not authenticate")
