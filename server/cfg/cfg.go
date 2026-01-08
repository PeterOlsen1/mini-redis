package cfg

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ConfigType struct {
	Server ServerConfig `yaml:"server"`
	Info   InfoConfig   `yaml:"info"`
	Log    LogConfig    `yaml:"log"`
}

type ServerConfig struct {
	// The address on which to run the server
	Address string `yaml:"address"`

	// The port on which to run the server
	Port int `yaml:"port"`

	// The number of ms to check and wipe expired TTLs
	// Set to <= 0 for no checking
	TTLCheck int `yaml:"ttl_check"`

	// Set to true if you want to require authentication
	RequireAuth bool `yaml:"require_auth"`

	// Define a lsit of users
	Users []User `yaml:"users"`

	// To be loaded once the application runs. Not defined in yaml
	LoadedUsers []User
}

// For basic operations, disabling logging will result in a ~17% performance increase
type InfoConfig struct {
	// Should we be collecting info?
	CollectOps bool `yaml:"collect_ops"`

	// Should we collect info on every command?
	// Adds overhead to every call instead of few
	Command bool `yaml:"command"`
}

type LogConfig struct {
	// If set to true, config will be sent to a file instead of STDIN
	File bool `yaml:"file"`

	// Log connection events
	Connect bool `yaml:"connect"`

	// Log disconnect events
	Disconnect bool `yaml:"disconnect"`

	// Log that data is sent by users
	DataEvent bool `yaml:"data_event"`

	// Log commands sent by users
	Command bool `yaml:"command"`
}

var config ConfigType
var Server ServerConfig
var Info InfoConfig
var Log LogConfig

var defaultConfig = ConfigType{
	Server: ServerConfig{
		Address:  "localhost",
		Port:     6379,
		TTLCheck: 2000,
	},
	Info: InfoConfig{
		CollectOps: true,
		Command:    true,
	},
	Log: LogConfig{
		File:       false,
		Connect:    true,
		Disconnect: true,
		DataEvent:  false,
		Command:    false,
	},
}

func LoadConfig(path string) error {
	f, err := os.Open(path)
	if err != nil { // just do the default if file open does not work
		fmt.Println("Failed to read config file, using defaults")
		config = defaultConfig
		Server = config.Server
		Info = config.Info
		Log = config.Log
		return nil
	}
	defer f.Close()

	// set default config before
	config = defaultConfig
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return err
	}

	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	fmt.Println("Error getting home directory:", err)
	// 	os.Exit(1)
	// }

	// homeFolder := filepath.Join(homeDir, ".mini-redis")
	// usersFilePath := filepath.Join(homeFolder, "users.acl")

	// err = os.MkdirAll(homeFolder, 0755)
	// if err != nil {
	// 	fmt.Println("Failed to create .mini-redis directory:", err)
	// 	os.Exit(1)
	// }

	// usersFile, err := os.OpenFile(usersFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	// if err != nil {
	// 	fmt.Println("Failed to open or create history file:", err)
	// 	os.Exit(1)
	// }

	// update individual config objects
	Server = config.Server
	Info = config.Info
	Log = config.Log

	if Server.RequireAuth && len(Server.Users) == 0 {
		return fmt.Errorf("must have one defined user if authentication is required")
	}

	return nil
}
