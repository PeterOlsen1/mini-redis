package errors

import (
	"fmt"
	"log"
	"mini-redis/server/auth/authtypes"
	"mini-redis/types/commands"
)

var GENERAL = fmt.Errorf("an error occoured")
var WRONGTYPE = fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
var NOT_INTEGER = fmt.Errorf("value is not an integer or out of range")
var INVALID_ARG = fmt.Errorf("INVALID_ARG failed to parse integer argument") //make into a function to return invaid argument?
func ARG_COUNT(cmd commands.Command, count int) error {
	log.Printf("ARG_COUNT error: %s", cmd.String())
	return fmt.Errorf("%s requires %d arguments", cmd, count)
}

func PERMS_GENERAL(cmd commands.Command) error {
	log.Printf("PERMS_GENERAL error: %s", cmd.String())
	return fmt.Errorf("you do not have permissions to run %s", cmd.String())
}

func PERMS_KEY(cmd commands.Command, perm authtypes.UserPermission, key string) error {
	log.Printf("PERMS_KEY %s on %s requires %s privileges", cmd.String(), key, perm.String())
	return fmt.Errorf("%s on %s requires %s privileges", cmd.String(), key, perm.String())
}

func PERMISSIONS(cmd commands.Command, perm authtypes.UserPermission) error {
	log.Printf("%s requies %s privileges", cmd.String(), perm.String())
	return fmt.Errorf("%s requies %s privileges", cmd.String(), perm.String())
}

var ALREADY_AUTH = fmt.Errorf("user is already authenticated. use logout to destroy session")
var COULD_NOT_AUTHENTICATE = fmt.Errorf("could not authenticate")
var USER_EXISTS = fmt.Errorf("user already exists")
