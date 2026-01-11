package auth

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
)

func OpenACLFile() (*os.File, error) {
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

	userFile, err := os.OpenFile(usersFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	return userFile, nil
}

func GetACLUsers() ([]User, error) {
	userFile, err := OpenACLFile()
	if err != nil {
		return nil, err
	}

	defer userFile.Close()

	decoder := yaml.NewDecoder(userFile)
	users := make([]User, 0)
	err = decoder.Decode(&users)
	if err != nil {
		if err == io.EOF {
			return users, nil
		}
		return nil, err
	}

	return users, nil
}

func AddACLUser(username string, password string, perms int) error {
	userFile, err := OpenACLFile()
	if err != nil {
		return err
	}

	defer userFile.Close()

	decoder := yaml.NewDecoder(userFile)
	users := make([]User, 0)
	err = decoder.Decode(&users)
	if err != nil {
		return err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := User{
		Username: username,
		Password: string(hashedPass),
		Perms:    perms,
	}

	users = append(users, newUser)

	encoder := yaml.NewEncoder(userFile)
	return encoder.Encode(users)
}

func RemoveACLUser(username string) error {
	userFile, err := OpenACLFile()
	if err != nil {
		return err
	}

	defer userFile.Close()

	decoder := yaml.NewDecoder(userFile)
	users := make([]User, 0)
	err = decoder.Decode(&users)
	if err != nil {
		return err
	}

	found := false
	for i, u := range users {
		if u.Username == username {
			users = append(users[:i], users[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("user could not be found")
	}

	encoder := yaml.NewEncoder(userFile)
	return encoder.Encode(users)
}

func CheckACLUser(username string, password string) (*User, error) {
	userFile, err := OpenACLFile()
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
