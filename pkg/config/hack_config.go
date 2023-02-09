package config

var (
	TestConfig = Config{
		ES: ESConf{
			Url: []string{
				"http://127.0.0.1:9200",
			},
		},
	}
)
