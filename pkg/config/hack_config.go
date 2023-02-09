package config

import (
	"os"
)

var (
	TestConfig = Config{
		Mysql: MySqlConf{
			HOST:   "mysql",
			DB:     "test",
			USER:   "root",
			PASSWD: "root",
		},
		Redis: RedisConf{
			Endpoint: "redis:6379",
		},
		ES: ESConf{
			Url: []string{
				"http://127.0.0.1:9200",
			},
		},
	}
)

func init() {
	if os.Getenv("CI") == "" {
		TestConfig.Mysql.HOST = "127.0.0.1"
		TestConfig.Redis.Endpoint = "127.0.0.1:6379"
	}
}
