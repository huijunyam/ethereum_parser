package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	RedidHost  string `yaml:"RedisHost"`
	RedisPort  string `yaml:"RedisPort"`
	ServerPort string `yaml:"ServerPort"`
}

var Conf *Config

func Init() {
	confBuf, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		panic(err)
	}
	var config Config
	if err := yaml.Unmarshal(confBuf, &config); err != nil {
		panic(err)
	}
	Conf = &config
}
