package cfg

// UPDATE THIS LATER TO TAKE A YAML CONFIG ARGUMENT ON SERVER START

type ConfigType struct {
	Server ServerConfig
	Info   InfoConfig
}

type ServerConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type InfoConfig struct {
	LogCommands bool `yaml:"log_commands"`
}

var Config ConfigType

func InitConfig() error {
	// do some file reading here
	Config.Server.Address = "localhost"
	Config.Server.Port = 6379
	Config.Info.LogCommands = true

	return nil
}
