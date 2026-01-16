package auth

import (
	"fmt"
	"io"
	"mini-redis/server/auth/authtypes"
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

func UpdateACLFile() error {
	userFile, err := OpenACLFile(true)
	if err != nil {
		return err
	}
	defer userFile.Close()

	encoder := yaml.NewEncoder(userFile)
	err = encoder.Encode(GetAllUsers())
	if err != nil {
		return err
	}

	return nil
}

// Opens the ACL file and reads all users initially
func LoadACLUsers() error {
	userFile, err := OpenACLFile(false)
	if err != nil {
		return err
	}

	defer userFile.Close()

	decoder := yaml.NewDecoder(userFile)
	users := make([]authtypes.User, 0)
	err = decoder.Decode(&users)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	for _, user := range users {
		SetUser(&user)
	}

	return nil
}

// Add a user to the ACL file
//
// Note for myslef:
// Go uses double pointers to signify if the function changes the pointer argument.
// Only use if you want to reassign the callee's pointer
func AddACLUser(user *authtypes.User, username string, password string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &authtypes.User{
		Username: username,
		Password: string(hashedPass),
		Perms:    0,
		Rules:    make(authtypes.Ruleset, 0),
	}

	SetUser(newUser)

	return UpdateACLFile()
}

func RemoveACLUser(username string) error {
	DeleteUser(username)
	return UpdateACLFile()
}

func CheckACLUser(currentUser **authtypes.User, username string, password string) error {
	for _, u := range GetAllUsers() {
		if u.Username != username {
			continue
		}

		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
		if err == nil {
			*currentUser = u
			return nil
		}
	}

	return fmt.Errorf("user not found")
}
