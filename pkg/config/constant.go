package config

import "time"

const (
	DefaultPageSize = 1000

	DefaultLockDuration = 10 * time.Second

	ES_ENDPOINTS = "ES_ENDPOINTS"
	ES_USERNAME  = "ES_USERNAME"
	ES_PASSWORD  = "ES_PASSWORD"

	ES_Mapping = `{
		"mappings": {
		  "properties": {
			"word":    { "type": "keyword" },  
			"en":    {"type": "text"},
			"cn":    {"type": "text"},
			"examples":  {"type": "text"}, 
			"synomyms":  {"type": "text"}, 
			"hits" : {"type": "date"}
		  }
		}
	  }`

	CibaIndexName = "ciba_translate"
)
