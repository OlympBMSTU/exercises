package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// config - one file

type Config struct {
	fileStorageDir string
	listenerPort   string
	db_host        string
	db_port        string
	database       string
	user           string
	password       string
}

func (cfg *Config) GetFileStorageName() string {
	return cfg.fileStorageDir
}

func (cfg *Config) GetHost() string {
	return cfg.db_host
}

func (cfg *Config) GetPort() string {
	return cfg.db_host
}

func (cfg *Config) GetDatabase() string {
	return cfg.database
}

func (cfg *Config) GetUser() string {
	return cfg.user
}

func (cfg *Config) GetPassword() string {
	return cfg.password
}

// it works but need to get path to dir
// error handling, maybe return struct string, err
// check

func Init() (*Config, error) {
	iniPath := "/etc/serv"

	file, err := os.Open(iniPath)

	fbytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Cant start server without initial file\n" +
			"Please creaate init file and put it to /etc/serv/")
		return nil, err
	}

	fileData := string(fbytes)
	configs := strings.Split(fileData, "\n")

	return &Config{
		configs[0],
		configs[1],
		configs[2],
		configs[3],
		configs[4],
		configs[5],
		configs[6],
	}, nil
}

var configInstance *Config = nil

func GetConfigInstance() (*Config, error) {
	if configInstance != nil {
		return configInstance, nil
	}

	var err error
	configInstance, err = Init()
	if err != nil {
		return nil, err
	}
	fmt.Print(configInstance.GetFileStorageName())
	return configInstance, nil
}
