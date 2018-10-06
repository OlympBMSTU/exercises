package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type PgConfig struct {
	db_host  string
	db_port  string
	database string
	user     string
	password string
}

func (cfg *PgConfig) GetHost() string {
	return cfg.db_host
}

func (cfg *PgConfig) GetPort() string {
	return cfg.db_host
}

func (cfg *PgConfig) GetDatabase() string {
	return cfg.database
}

func (cfg *PgConfig) GetUser() string {
	return cfg.user
}

func (cfg *PgConfig) GetPassword() string {
	return cfg.password
}

// it works but need to get path to dir
// error handling, maybe return struct string, err
// check

func InitPg() (*PgConfig, error) {
	iniPath := "/home/mavr/pg_conf" //"/etc/serv"

	file, err := os.Open(iniPath)

	fbytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Cant start server without initial file\n" +
			"Please creaate init file and put it to /etc/serv/")
		return nil, err
	}

	fileData := string(fbytes)
	configs := strings.Split(fileData, "\n")

	return &PgConfig{
		configs[0],
		configs[1],
		configs[2],
		configs[3],
		configs[4]}, nil
}

var pgConfigInstance *PgConfig = nil

func GetPgConfigInstance() (*PgConfig, error) {
	if pgConfigInstance != nil {
		return pgConfigInstance, nil
	}

	var err error
	pgConfigInstance, err = InitPg()
	if err != nil {
		return nil, err
	}
	// fmt.Print(pgConfigInstance.GetFileStorageName())
	return pgConfigInstance, nil
}
