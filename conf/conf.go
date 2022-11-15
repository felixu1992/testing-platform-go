package conf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	config2 "testing-platform-go/internal/config"
)

var Server config2.ServerConf
var Application config2.ApplicationConf
var Datasource config2.DatasourceConf
var Redis config2.RedisConf

func InitConf() {
	config, err := ioutil.ReadFile("application.yml")
	if err != nil {
		return
	}
	var conf Config
	if err := yaml.Unmarshal(config, &conf); err != nil {
		return
	}
	Server = conf.Server
	Application = conf.Application
	Datasource = conf.Datasource
	Redis = conf.Redis
}

type Config struct {
	Server      config2.ServerConf      `yaml:"server"`
	Application config2.ApplicationConf `yaml:"application"`
	Datasource  config2.DatasourceConf  `yaml:"datasource"`
	Redis       config2.RedisConf       `yaml:"redis"`
}
