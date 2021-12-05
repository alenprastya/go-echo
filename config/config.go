package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAMER    string
}

func GetConfirations() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}
