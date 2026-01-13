package auth

import (
	"fmt"
	"io"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/cfg"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
)

func OpenACLFile(truncate bool) (*os.File, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}

	homeFolder := filepath.Join(homeDir, ".mini-redis")
	usersFilePath := filepath.Join(homeFolder, "users.acl")

	err = os.MkdirAll(homeFolder, 0755)
	if err != nil {
		return nil, err
	}

	flags := os.O_CREATE | os.O_RDWR
	if truncate {
		flags |= os.O_TRUNC
	} else {
		flags |= os.O_APPEND
	}

	userFile, err := os.OpenFile(usersFilePath, flags, 0644)
	if err != nil {
		return nil, err
	}

	return userFile, nil
}

// Opens the ACL file and reads all users initially
func LoadACLUsers() ([]authtypes.User, error) {
	userFile, err := OpenACLFile(false)
	if err != nil {
		return nil, err
	}

	defer userFile.Close()

	decoder := yaml.NewDecoder(userFile)
	users := make([]authtypes.User, 0)
	err = decoder.Decode(&users)
	if err != nil {
		if err == io.EOF {
			return users, nil
		}
		return nil, err
	}

	return users, nil
}

// Add a user to the ACL file
func AddACLUser(username string, password string, perms int, rules authtypes.Ruleset) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := authtypes.User{
		Username: username,
		Password: string(hashedPass),
		Perms:    perms,
		Rules:    rules,
	}

	cfg.Server.LoadedUsers = append(cfg.Server.LoadedUsers, newUser)

	userFile, err := OpenACLFile(true)
	if err != nil {
		return err
	}
	defer userFile.Close()

	encoder := yaml.NewEncoder(userFile)
	err = encoder.Encode(cfg.Server.LoadedUsers)
	if err != nil {
		return err
	}

	return nil
}

func RemoveACLUser(username string) error {
	found := false
	users := cfg.Server.LoadedUsers

	for i, u := range users {
		if u.Username == username {
			cfg.Server.LoadedUsers = append(users[:i], users[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("user could not be found")
	}

	// Reopen the file in truncate mode for writing
	userFile, err := OpenACLFile(true)
	if err != nil {
		return err
	}
	defer userFile.Close()

	encoder := yaml.NewEncoder(userFile)
	err = encoder.Encode(users)
	if err != nil {
		return err
	}

	return nil
}

func CheckACLUser(username string, password string) (*authtypes.User, error) {
	userFile, err := OpenACLFile(false)
	if err != nil {
		return nil, err
	}

	defer userFile.Close()

	decoder := yaml.NewDecoder(userFile)
	users := make([]authtypes.User, 0)
	err = decoder.Decode(&users)
	if err != nil && err != io.EOF {
		return nil, err
	}

	for _, u := range users {
		if u.Username != username {
			continue
		}

		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
		if err == nil {
			return &u, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}
