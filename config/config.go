package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

// config - one file

type Config struct {
	fileStorageDir string
	listenerPort   string
	dbHost         string
	dbPort         string
	database       string
	dbUser         string
	dbPassword     string
	smtpHost       string
	smtpPort       string
	smtpUser       string
	smtpPassword   string
	acceptorMail   string
	mailSubject    string
}

func (cfg Config) GetFileStorageName() string {
	return cfg.fileStorageDir
}

func (cfg Config) GetDBHost() string {
	return cfg.dbHost
}

func (cfg Config) GetDBPort() string {
	return cfg.dbPort
}

func (cfg *Config) GetDatabase() string {
	return cfg.database
}

func (cfg Config) GetDBUser() string {
	return cfg.dbUser
}

func (cfg Config) GetDBPassword() string {
	return cfg.dbPassword
}

func (cfg Config) GetSMTPHost() string {
	return cfg.smtpHost
}

func (cfg Config) GetSMTPPort() string {
	return cfg.smtpPort
}

func (cfg Config) GetSMTPUser() string {
	return cfg.smtpUser
}

func (cfg Config) GetSMTPPassword() string {
	return cfg.smtpPassword
}

func (cfg Config) GetAcceptorMail() string {
	return cfg.acceptorMail
}

func (cfg Config) GetMailSubject() string {
	return cfg.mailSubject
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

	countFields := reflect.ValueOf(Config{}).NumField()
	if len(configs) < countFields {
		fmt.Println("Not enough fields")
		return nil, err
	}

	return &Config{
		configs[0],
		configs[1],
		configs[2],
		configs[3],
		configs[4],
		configs[5],
		configs[6],
		configs[7],
		configs[8],
		configs[9],
		configs[10],
		configs[11],
		configs[12],
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
