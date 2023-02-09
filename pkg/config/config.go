package config

import (
	"encoding/json"
	"io/ioutil"
)

type MySqlConf struct {
	HOST   string `json:"host,omitempty"`
	DB     string `json:"db,omitempty"`
	USER   string `json:"user,omitempty"`
	PASSWD string `json:"passwd,omitempty"`
}

type RedisConf struct {
	Endpoint string `json:"endpoint,omitempty"`
	Password string `json:"password,omitempty"`
}

type Config struct {
	EnableProfile bool      `json:"enable_profile"`
	Debug         bool      `json:"debug,omitempty"`
	Mysql         MySqlConf `json:"mysql,omitempty"`
	TplFolder     string    `json:"tpl_folder,omitempty"`
	Redis         RedisConf `json:"redis,omitempty"`
	ES            ESConf    `json:"es,omitempty"`
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
