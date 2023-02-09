package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	EnableProfile bool   `json:"enable_profile"`
	Debug         bool   `json:"debug,omitempty"`
	ES            ESConf `json:"es,omitempty"`
}

func LoadConfigFromFile(configFileName string, o *Config) error {
	if err := loadConfig(configFileName, o); err != nil {
		return err
	}
	return nil
}

func loadConfig(configFileName string, o interface{}) error {
	bytesBody, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytesBody, o); err != nil {
		return err
	}
	return nil
}

type ESConf struct {
	Url      []string `json:"url,omitempty"`
	Username string   `json:"username,omitempty"`
	Password string   `json:"password,omitempty"`
}
