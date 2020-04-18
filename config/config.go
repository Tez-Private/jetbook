package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// EnvConifg use Environment variables
var EnvConfig *ConfigList

// ConfigList  is Setting
type ConfigList struct {
	// MySQL Config List
	MYSQL_USER   string
	MYSQL_PASS   string
	MYSQL_DBNAME string
	MYSQL_HOST   string
	MYSQL_ACCESS string
}

// ReadLocalEnvConfig is For test
func ReadLocalEnvConfig() {
	conf := ConfigList{}
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	EnvConfig = &conf
}
