package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type configuration struct {
	Database database
}

type database struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

var Config configuration

func loadConfig() error {
	configPath := os.Getenv("GO_CONFIG_PATH")
	if configPath == "" {
		return fmt.Errorf("Can't find GO_CONFIG_PATH env variable.")
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &Config)
}

func init() {
	if err := loadConfig(); err != nil {
		log.Fatal("Could not load config: ", err)
	}
}
