package config

type DatasourceConf struct {
	Ip       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
