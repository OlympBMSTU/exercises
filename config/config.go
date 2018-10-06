package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	fileStorageDir string
	listenerPort   string
	dbPath         string
}

func (cfg *Config) GetFileStorageName() string {
	return cfg.fileStorageDir
}

// it works but need to get path to dir
// error handling, maybe return struct string, err
// check

func Init() (*Config, error) {
	iniPath := "/home/mavr/conf" //"/etc/serv"

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
